package main

import (
	"momentum-core/config"
	"momentum-core/routers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "momentum-core/docs"
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
	dispatcher.valueRouter.RegisterValueRoutes(dispatcher.server)

	dispatcher.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return dispatcher
}

func (d *Dispatcher) Serve() {
	d.server.Run("localhost:8080")
}
