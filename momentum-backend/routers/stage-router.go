package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ROUTING_PATH_STAGE_BY_ID = VERSION + "/repository/:repositoryName/app/stage/:stageId"
const ROUTING_PATH_STAGE = VERSION + "/stage"

type StageRouter struct {
	stageService *services.StageService
}

func NewStageRouter(stageService *services.StageService) *StageRouter {

	router := new(StageRouter)

	router.stageService = stageService

	return router
}

func (s *StageRouter) RegisterStageRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_STAGE_BY_ID, s.getStage)
	server.POST(ROUTING_PATH_STAGE, s.addStage)
}

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

func (sr *StageRouter) addStage(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	request, err := models.ExtractStageCreateRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	stage, err := sr.stageService.AddStage(request, traceId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.IndentedJSON(http.StatusOK, stage)
}
