package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	gittransaction "github.com/Joel-Haeberli/git-transaction"
	"github.com/gin-gonic/gin"
)

const ROUTING_PATH_STAGE_BY_ID = VERSION + "/repository/:repositoryName/app/stage/:stageId"
const ROUTING_PATH_STAGE = VERSION + "/stage"

type StageRouter struct {
	stageService      *services.StageService
	repositoryService *services.RepositoryService
	config            *config.MomentumConfig
}

func NewStageRouter(stageService *services.StageService, repositoryService *services.RepositoryService, config *config.MomentumConfig) *StageRouter {

	router := new(StageRouter)

	router.stageService = stageService
	router.repositoryService = repositoryService
	router.config = config

	return router
}

func (s *StageRouter) RegisterStageRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_STAGE_BY_ID, s.getStage)
	server.POST(ROUTING_PATH_STAGE, s.addStage)
}

// GetStage godoc
//
//	@Summary		get a stage of a repository by id
//	@Tags			stages
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Param			stageId				path		string					true	"Stage ID"
//	@Success		200		{object}	models.Deployment
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/app/stage/{stageId} [get]
func (s *StageRouter) getStage(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	stageId := c.Param("stageId")

	result, err := s.stageService.GetStage(repoName, stageId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// AddStage godoc
//
//	@Summary		add a new stage
//	@Tags			stages
//	@Accept			json
//	@Produce		json
//	@Param			stageCreateRequest	body		models.StageCreateRequest	true	"Create Stage"
//	@Success		200		{object}	models.Stage
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/stage [post]
func (sr *StageRouter) addStage(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	request, err := models.ExtractStageCreateRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	repo, err := sr.repositoryService.GetRepository(request.RepositoryName, traceId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	ctx, transaction, err := gittransaction.New(TRANSACTION_MODE, repo.Path, sr.config.TransactionToken())

	stage, err := sr.stageService.AddStage(request, traceId)
	if err != nil {
		transaction.Rollback(ctx)
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Write(ctx)
	if err != nil {
		transaction.Rollback(ctx)
		c.IndentedJSON(http.StatusInternalServerError, models.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Commit(ctx)
	if err != nil {
		transaction.Rollback(ctx)
		c.IndentedJSON(http.StatusInternalServerError, models.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.IndentedJSON(http.StatusOK, stage)
}
