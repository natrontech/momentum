package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ROUTING_PATH_VALUE_BY_ID = VERSION + "/repository/:repositoryName/value/:valueId"
const ROUTING_PATH_VALUE = VERSION + "/value"
const ROUTING_PATH_VALUE_BY_APPLICATION = VERSION + "/repository/:repositoryName/application/values/:applicationId"
const ROUTING_PATH_VALUE_BY_STAGE = VERSION + "/repository/:repositoryName/stage/values/:stageId"
const ROUTING_PATH_VALUE_BY_DEPLOYMENT = VERSION + "/repository/:repositoryName/deployment/values/:deploymentId"

type ValueRouter struct {
	valueService *services.ValueService
}

func NewValueRouter(valueService *services.ValueService) *ValueRouter {

	vs := new(ValueRouter)

	vs.valueService = valueService

	return vs
}

func (vr *ValueRouter) RegisterValueRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_VALUE_BY_ID, vr.valueById)
	server.GET(ROUTING_PATH_VALUE_BY_APPLICATION, vr.valuesByApplication)
	server.GET(ROUTING_PATH_VALUE_BY_STAGE, vr.valuesByStage)
	server.GET(ROUTING_PATH_VALUE_BY_DEPLOYMENT, vr.valuesByDeployment)
}

// valueById godoc
//
//	@Summary		get a value of a repository by id
//	@Tags			values
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Param			valueId				path		string					true	"Value ID"
//	@Success		200		{object}	models.Value
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/value/{valueId} [get]
func (vr *ValueRouter) valueById(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	valueId := c.Param("valueId")

	result, err := vr.valueService.ValueById(repoName, valueId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// valuesByApplication godoc
//
//	@Summary		get all values of an application by  the applications id
//	@Tags			values
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Param			applicationId		path		string					true	"Application ID"
//	@Success		200		{array}		models.ValueWrapper
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/application/values/{applicationId} [get]
func (vr *ValueRouter) valuesByApplication(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	applicationId := c.Param("applicationId")

	result, err := vr.valueService.ValuesByApplication(repoName, applicationId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// valuesByStage godoc
//
//	@Summary		get all values of an stage by  the stages id
//	@Tags			values
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Param			stageId				path		string					true	"Stage ID"
//	@Success		200		{array}		models.ValueWrapper
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/stage/values/{stageId} [get]
func (vr *ValueRouter) valuesByStage(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	stageId := c.Param("stageId")

	result, err := vr.valueService.ValuesByStage(repoName, stageId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// valuesByDeployment godoc
//
//	@Summary		get all values of a deployment by the deployments id
//	@Tags			values
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Param			deploymentId		path		string					true	"Deployment ID"
//	@Success		200		{array}		models.ValueWrapper
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/deployment/values/{deploymentId} [get]
func (vr *ValueRouter) valuesByDeployment(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	deploymentId := c.Param("deploymentId")

	result, err := vr.valueService.ValuesByDeployment(repoName, deploymentId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}
