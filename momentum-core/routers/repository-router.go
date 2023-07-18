package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const VERSION = "" // for example "/v1" or "/beta"

const ROUTING_PATH_REPOSITORIES = VERSION + "/repositories"
const ROUTING_PATH_REPOSITORY = VERSION + "/repository"

const ROUTING_PATH_REPOSITORY_APPLICATIONS = VERSION + "/repository/:repositoryName/applications"
const ROUTING_PATH_REPOSITORY_STAGES = VERSION + "/repository/:repositoryName/stages"
const ROUTING_PATH_REPOSITORY_DEPLOYMENTS = VERSION + "/repository/:repositoryName/deployments"

const ROUTING_PATH_REPOSITORY_BY_NAME = VERSION + "/repository/:repositoryName/"

type RepositoryRouter struct {
	repositoryService  *services.RepositoryService
	applicationService *services.ApplicationService
	stageService       *services.StageService
	deploymentService  *services.DeploymentService
}

func NewRepositoryRouter(repositoryService *services.RepositoryService,
	applicationService *services.ApplicationService,
	stageService *services.StageService,
	deploymentService *services.DeploymentService,
) *RepositoryRouter {

	router := new(RepositoryRouter)

	router.repositoryService = repositoryService
	router.applicationService = applicationService
	router.stageService = stageService
	router.deploymentService = deploymentService

	return router
}

func (r *RepositoryRouter) RegisterRepositoryRoutes(server *gin.Engine) {

	server.POST(ROUTING_PATH_REPOSITORY, r.addRepository)
	server.GET(ROUTING_PATH_REPOSITORIES, r.getRepositories)
	server.GET(ROUTING_PATH_REPOSITORY_BY_NAME, r.getRepository)
	server.GET(ROUTING_PATH_REPOSITORY_APPLICATIONS, r.getApplications)
	server.GET(ROUTING_PATH_REPOSITORY_STAGES, r.getStages)
	server.GET(ROUTING_PATH_REPOSITORY_DEPLOYMENTS, r.getDeployments)
}

// GetRepositories godoc
//
//	@Summary		load repositories
//	@Description	load all repositories managed by this instance
//	@Tags			repositories
//	@Produce		json
//	@Success		200		{array}		models.Repository
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repositories [get]
func (r *RepositoryRouter) getRepositories(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	repos := r.repositoryService.GetRepositories(traceId)
	c.JSON(http.StatusOK, repos)
}

// AddRepository godoc
//
//	@Summary		add a new repository
//	@Description	adds a new repository to the instance
//	@Tags			repositories
//	@Accept			json
//	@Produce		json
//	@Param			repositoryCreateRequest	body		models.RepositoryCreateRequest	true	"Create Repository"
//	@Success		200		{object}	models.Repository
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository [post]
func (r *RepositoryRouter) addRepository(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	request, err := models.ExtractRepositoryCreateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	repo, err := r.repositoryService.AddRepository(request, traceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, repo)
}

type repoUrl struct {
	repositoryName string `uri:"repositoryName" binding:"required"`
}

// GetRepository godoc
//
//	@Summary		get a repository
//	@Tags			repositories
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Success		200		{object}	models.Repository
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName} [get]
func (r *RepositoryRouter) getRepository(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	req := new(repoUrl)
	err := c.BindUri(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	result, err := r.repositoryService.GetRepository(req.repositoryName, traceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetApplications godoc
//
//	@Summary		get all applications of a repository
//	@Tags			applications
//	@Accept			json
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Success		200		{array}		models.Application
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/applications [get]
func (r *RepositoryRouter) getApplications(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	req := new(repoUrl)
	err := c.BindUri(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	result, err := r.applicationService.GetApplications(req.repositoryName, traceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetStages godoc
//
//	@Summary		get stages
//	@Tags			stages
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Success		200		{array}		models.Stage
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/stages [get]
func (r *RepositoryRouter) getStages(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	req := new(repoUrl)
	err := c.BindUri(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	result, err := r.stageService.GetStages(req.repositoryName, traceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetDeployments godoc
//
//	@Summary		get deployments
//	@Tags			deployments
//	@Produce		json
//	@Param			repositoryName		path		string					true	"Repository Name"
//	@Success		200		{array}		models.Deployment
//	@Failure		400		{object}	models.ApiError
//	@Failure		404		{object}	models.ApiError
//	@Failure		500		{object}	models.ApiError
//	@Router			/repository/{repositoryName}/deployments [get]
func (r *RepositoryRouter) getDeployments(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	req := new(repoUrl)
	err := c.BindUri(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	result, err := r.deploymentService.GetDeployments(req.repositoryName, traceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}
