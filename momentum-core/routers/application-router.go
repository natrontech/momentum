package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	gittransaction "github.com/Joel-Haeberli/git-transaction"
	"github.com/gin-gonic/gin"
)

const ROUTING_PATH_APPLICATION_BY_ID = VERSION + "/repository/:repositoryName/:applicationId"
const ROUTING_PATH_APPLICATION = VERSION + "/application"

type ApplicationRouter struct {
	applicationService *services.ApplicationService
	repositoryService  *services.RepositoryService
	config             *config.MomentumConfig
}

func NewApplicationRouter(applicationService *services.ApplicationService,
	repositoryService *services.RepositoryService,
	config *config.MomentumConfig,
) *ApplicationRouter {

	router := new(ApplicationRouter)

	router.applicationService = applicationService
	router.repositoryService = repositoryService
	router.config = config

	return router
}

func (a *ApplicationRouter) RegisterApplicationRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_APPLICATION_BY_ID, a.getApplication)
	server.POST(ROUTING_PATH_APPLICATION, a.addApplication)
}

// GetApplication godoc
//
//	@Summary		get an application of a repository by id
//	@Tags			applications
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Param			applicationId		path		string					true	"Application ID"
//	@Success		200		{object}	models.Application
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/{applicationId} [get]
func (a *ApplicationRouter) getApplication(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	// TODO: find out why binding doesnt work:
	// type AppUrl struct {
	// 	repositoryName string `uri:"repositoryName" binding:"required"`
	// 	applicationId  string `uri:"applicationId" binding:"required"`
	// }
	// var url AppUrl
	// err := c.BindUri(&url)

	// fmt.Println(url.repositoryName, url.applicationId)

	// if err != nil {
	// 	fmt.Println("error after binding:", err.Error())
	// 	c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
	// 	return
	// }
	// fmt.Println("application-id:", url.applicationId, "/ by hand:", c.Request.RequestURI, url)

	repoName := c.Param("repositoryName")
	appId := c.Param("applicationId")

	result, err := a.applicationService.GetApplication(repoName, appId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// AddApplication godoc
//
//	@Summary		add an application
//	@Tags			applications
//	@Accept			json
//	@Produce		json
//	@Param			applicationCreateRequest	body		models.ApplicationCreateRequest	true	"Create Application"
//	@Success		200		{object}	models.Application
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/application [post]
func (a *ApplicationRouter) addApplication(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	request, err := models.ExtractApplicationCreateRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	repo, err := a.repositoryService.GetRepository(request.RepositoryName, traceId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	ctx, transaction, err := gittransaction.New(TRANSACTION_MODE, repo.Path, a.config.TransactionToken())

	application, err := a.applicationService.AddApplication(request, traceId)
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

	c.IndentedJSON(http.StatusOK, application)
}
