package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	gittransaction "github.com/Joel-Haeberli/git-transaction"
	"github.com/gin-gonic/gin"
)

const ROUTING_PATH_DEPLOYMENT_BY_ID = VERSION + "/repository/:repositoryName/app/stage/deployment/:deploymentId"
const ROUTING_PATH_DEPLOYMENT = VERSION + "/deployment"

type DeploymentRouter struct {
	deploymentService *services.DeploymentService
	repositoryService *services.RepositoryService
	config            *config.MomentumConfig
}

func NewDeploymentRouter(deploymentService *services.DeploymentService, repositoryService *services.RepositoryService, config *config.MomentumConfig) *DeploymentRouter {

	router := new(DeploymentRouter)

	router.deploymentService = deploymentService
	router.repositoryService = repositoryService
	router.config = config

	return router
}

func (d *DeploymentRouter) RegisterDeploymentRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_DEPLOYMENT_BY_ID, d.getDeployment)
	server.POST(ROUTING_PATH_DEPLOYMENT, d.addDeployment)
}

// GetDeployment godoc
//
//	@Summary		get a deployment of a repository by id
//	@Tags			deployments
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Param			deploymentId		path		string					true	"Deployment ID"
//	@Success		200		{object}	models.Deployment
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/app/stage/deployment/{deploymentId} [get]
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

// GetDeployment godoc
//
//	@Summary		get a deployment of a repository by id
//	@Tags			deployments
//	@Accept			json
//	@Produce		json
//	@Param			deploymentCreateRequest	body		models.DeploymentCreateRequest	true	"Create Deployment"
//	@Success		200		{object}	models.Deployment
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/deployment [post]
func (d *DeploymentRouter) addDeployment(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	request, err := models.ExtractDeploymentCreateRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	repo, err := d.repositoryService.GetRepository(request.RepositoryName, traceId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	ctx, transaction, err := gittransaction.New(gittransaction.SINGLEBRANCH, repo.Path, d.config.TransactionToken())

	deployment, err := d.deploymentService.AddDeployment(request, traceId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Write(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Commit(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.IndentedJSON(http.StatusOK, deployment)
}
