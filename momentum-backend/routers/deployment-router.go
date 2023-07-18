package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ROUTING_PATH_DEPLOYMENT_BY_ID = VERSION + "/repository/:repositoryName/app/stage/deployment/:deploymentId"
const ROUTING_PATH_DEPLOYMENT = VERSION + "/deployment"

type DeploymentRouter struct {
	deploymentService *services.DeploymentService
}

func NewDeploymentRouter(deploymentService *services.DeploymentService) *DeploymentRouter {

	router := new(DeploymentRouter)

	router.deploymentService = deploymentService

	return router
}

func (d *DeploymentRouter) RegisterDeploymentRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_DEPLOYMENT_BY_ID, d.getDeployment)
	server.POST(ROUTING_PATH_DEPLOYMENT, d.addDeployment)
}

func (d *DeploymentRouter) getDeployment(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	deploymentId := c.Param("deploymentId")

	result, err := d.deploymentService.GetDeployment(repoName, deploymentId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (d *DeploymentRouter) addDeployment(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	request, err := models.ExtractDeploymentCreateRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	deployment, err := d.deploymentService.AddDeployment(request, traceId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.IndentedJSON(http.StatusOK, deployment)
}
