package docs

import "momentum-core/models"

// swagger:route POST /application application add-app-routeid
//
// adds repository
//
// responses:
//   200: applicationResponse
//   400: errorResponse
//   500: errorResponse

// swagger:route GET /repository/{repositoryName}/{applicationId} repository get-app-routeid
//
// read an application from defined repository repositoryName and applicationId
//
// parameters:
//	+ name: repositoryName
//	  in: path
// 	  description: name of the repository
//	  required: true
// 	  type: string
//	+ name: applicationId
//	  in: path
// 	  description: id of the application
//	  required: true
// 	  type: string
//
// responses:
//   200: applicationResponse
//   400: errorResponse
//   500: errorResponse

// swagger:response applicationResponse
type applicationResponseWrapper struct {
	// in:body
	Body models.Application
}

// swagger:parameters add-app-routeid
type applicationCreationWrapper struct {
	// in:body
	Body models.ApplicationCreateRequest
}
