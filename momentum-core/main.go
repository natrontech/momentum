package main

import (
	"fmt"
	"momentum-core/clients"
	"momentum-core/config"
	"momentum-core/routers"
	"momentum-core/services"

	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

// @title		Momentum Core API
// @version		early-alpha
// @description	The momentum core api manages the core structure of momentum
//
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html
//
// @schemes 	http, https
// @host		localhost:8080
// @BasePath	/
func main() {

	fmt.Println("Starting momentum-core")

	config, err := config.InitializeMomentumCore()
	if err != nil {
		panic("failed initializing momentum. problem: " + err.Error())
	}

	gitClient := clients.NewGitClient(config)
	kustomizeClient := clients.NewKustomizationValidationClient(config)

	templateService := services.NewTemplateService()
	treeService := services.NewTreeService(config)
	repositoryService := services.NewRepositoryService(config, treeService, gitClient, kustomizeClient)
	applicationService := services.NewApplicationService(config, treeService, templateService)
	stageService := services.NewStageService(config, treeService, templateService)
	deploymentService := services.NewDeploymentService(config, stageService, templateService, treeService)
	valueService := services.NewValueService(treeService)

	templateRouter := routers.NewTemplateRouter()
	valueRouter := routers.NewValueRouter(valueService)
	deploymentRouter := routers.NewDeploymentRouter(deploymentService)
	stageRouter := routers.NewStageRouter(stageService)
	applicationRouter := routers.NewApplicationRouter(applicationService)
	repositoryRouter := routers.NewRepositoryRouter(repositoryService, applicationService, stageService, deploymentService)

	dispatcher := NewDispatcher(config, repositoryRouter, applicationRouter, stageRouter, deploymentRouter, valueRouter, templateRouter)

	dispatcher.Serve()
}
