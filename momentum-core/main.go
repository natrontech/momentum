package main

import (
	"fmt"
	"momentum-core/config"

	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

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

	fmt.Println("Starting momentum-core")

	config, err := config.InitializeMomentumCore()
	if err != nil {
		panic("failed initializing momentum. problem: " + err.Error())
	}

	// gitClient := clients.NewGitClient(config)
	// kustomizeClient := clients.NewKustomizationValidationClient(config)

	dispatcher := NewDispatcher(config)

	dispatcher.Serve()
}
