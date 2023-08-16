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
const ROUTING_PATH_VALUE_BY_APPLICATION = VERSION + "/:repositoryName/application/values/:applicationId"
const ROUTING_PATH_VALUE_BY_STAGE = VERSION + "/:repositoryName/stage/values/:stageId"
const ROUTING_PATH_VALUE_BY_DEPLOYMENT = VERSION + "/:repositoryName/deployment/values/:deploymentId"

type ValueRouter struct {
	valueService *services.ValueService
}

func NewValueRouter(valueService *services.ValueService) *ValueRouter {
	return new(ValueRouter)
}

func (vr *ValueRouter) RegisterValueRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_VALUE_BY_ID, vr.valueById)
	server.GET(ROUTING_PATH_VALUE_BY_APPLICATION, vr.valuesByApplication)
	server.GET(ROUTING_PATH_VALUE_BY_STAGE, vr.valuesByStage)
	server.GET(ROUTING_PATH_VALUE_BY_DEPLOYMENT, vr.valuesByDeployment)
}

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

func (vr *ValueRouter) valuesByApplication(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	applicationId := c.Param("applicationId")

	result, err := vr.valueService.ValueById(repoName, applicationId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

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
