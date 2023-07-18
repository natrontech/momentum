package main

import (
	"momentum-core/config"
	"momentum-core/routers"

	"github.com/gin-gonic/gin"
)

type Dispatcher struct {
	server            *gin.Engine
	config            *config.MomentumConfig
	repositoryRouter  *routers.RepositoryRouter
	applicationRouter *routers.ApplicationRouter
	stageRouter       *routers.StageRouter
	deploymentRouter  *routers.DeploymentRouter
	valueRouter       *routers.ValueRouter
	templateRouter    *routers.TemplateRouter
}

func NewDispatcher(config *config.MomentumConfig,
	repositoryRouter *routers.RepositoryRouter,
	applicationRouter *routers.ApplicationRouter,
	stageRouter *routers.StageRouter,
	deploymentRouter *routers.DeploymentRouter,
	valueRouter *routers.ValueRouter,
	templateRouter *routers.TemplateRouter) *Dispatcher {

	dispatcher := new(Dispatcher)

	dispatcher.server = gin.Default()

	dispatcher.config = config

	dispatcher.repositoryRouter = repositoryRouter
	dispatcher.applicationRouter = applicationRouter
	dispatcher.stageRouter = stageRouter
	dispatcher.deploymentRouter = deploymentRouter
	dispatcher.valueRouter = valueRouter

	dispatcher.repositoryRouter.RegisterRepositoryRoutes(dispatcher.server)
	dispatcher.applicationRouter.RegisterApplicationRoutes(dispatcher.server)
	dispatcher.stageRouter.RegisterStageRoutes(dispatcher.server)
	dispatcher.deploymentRouter.RegisterDeploymentRoutes(dispatcher.server)

	return dispatcher
}

func (d *Dispatcher) Serve() {
	d.server.Run("localhost:8080")
}
