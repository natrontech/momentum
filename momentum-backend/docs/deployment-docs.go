package docs

import "momentum-core/models"

// swagger:route POST /deployment deployment add-deployment-routeid
//
// adds deployment
//
// responses:
//   200: deploymentResponse
//   400: errorResponse
//   500: errorResponse

// swagger:route GET /repository/{repositoryName}/app/stage/deployment/{deploymentId} repository get-deployment-routeid
//
// read a deployment from defined repository repositoryName and deploymentId
//
// parameters:
//	+ name: repositoryName
//	  in: path
// 	  description: name of the repository
//	  required: true
// 	  type: string
//	+ name: deploymentId
//	  in: path
// 	  description: id of the deployment
//	  required: true
// 	  type: string
//
// responses:
//   200: deploymentResponse
//   400: errorResponse
//   500: errorResponse

// swagger:response deploymentResponse
type deploymentResponseWrapper struct {
	// in:body
	Body models.Deployment
}

// swagger:parameters add-deployment-routeid
type deploymentCreationWrapper struct {
	// in:body
	Body models.DeploymentCreateRequest
}
