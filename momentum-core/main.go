package main

import (
	"fmt"
	"momentum-core/artefacts"
	"momentum-core/config"
	"momentum-core/files"
	"momentum-core/templates"

	"github.com/gin-gonic/gin"

	// do not change order of the three following imports. It would break stuff.
	_ "momentum-core/docs" // This line is necessary for swagger to find your docs!

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const momentumCoreAsciiArt = "███╗   ███╗ ██████╗ ███╗   ███╗███████╗███╗   ██╗████████╗██╗   ██╗███╗   ███╗     ██████╗ ██████╗ ██████╗ ███████╗\n████╗ ████║██╔═══██╗████╗ ████║██╔════╝████╗  ██║╚══██╔══╝██║   ██║████╗ ████║    ██╔════╝██╔═══██╗██╔══██╗██╔════╝\n██╔████╔██║██║   ██║██╔████╔██║█████╗  ██╔██╗ ██║   ██║   ██║   ██║██╔████╔██║    ██║     ██║   ██║██████╔╝█████╗  \n██║╚██╔╝██║██║   ██║██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║   ██║   ██║██║╚██╔╝██║    ██║     ██║   ██║██╔══██╗██╔══╝  \n██║ ╚═╝ ██║╚██████╔╝██║ ╚═╝ ██║███████╗██║ ╚████║   ██║   ╚██████╔╝██║ ╚═╝ ██║    ╚██████╗╚██████╔╝██║  ██║███████╗\n╚═╝     ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝    ╚═════╝ ╚═╝     ╚═╝     ╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝"

// @title		Momentum Core API
// @version		early-alpha
// @description	The momentum core api manages the core structure of momentum
//
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html
//
// @schemes 	http
// @host		localhost:8080
// @BasePath	/
func main() {

	_, err := config.InitializeMomentumCore()
	if err != nil {
		panic("failed initializing momentum. problem: " + err.Error())
	}

	fmt.Println(momentumCoreAsciiArt)

	// gitClient := clients.NewGitClient(config)
	// kustomizeClient := clients.NewKustomizationValidationClient(config)

	//gin.SetMode(gin.ReleaseMode)

	server := gin.Default()

	files.RegisterFileRoutes(server)
	artefacts.RegisterArtefactRoutes(server)
	templates.RegisterTemplateRoutes(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run("localhost:8080")
}
