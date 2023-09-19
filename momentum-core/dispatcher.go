package main

import (
	"momentum-core/artefacts"
	"momentum-core/config"
	"momentum-core/files"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "momentum-core/docs"
)

type Dispatcher struct {
	server *gin.Engine
	config *config.MomentumConfig
}

func NewDispatcher(config *config.MomentumConfig) *Dispatcher {

	dispatcher := new(Dispatcher)
	dispatcher.config = config
	dispatcher.server = gin.Default()

	files.RegisterFileRoutes(dispatcher.server)
	artefacts.RegisterArtefactRoutes(dispatcher.server)

	dispatcher.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return dispatcher
}

func (d *Dispatcher) Serve() {
	d.server.Run("localhost:8080")
}
