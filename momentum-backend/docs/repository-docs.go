package docs

import "momentum-core/models"

// swagger:route POST /repository repository add-repo-routeid
//
// adds repository
//
// responses:
//   200: repositoryResponse
//   400: errorResponse
//   500: errorResponse

// swagger:route GET /repository/{repositoryName} repository get-repo-routeid
//
// read a repository with defined repositoryName
//
// parameters:
//   + name: repositoryName
//	  in: path
// 	  description: name of the repository
//	  required: true
// 	  type: string
//
// responses:
//   200: repositoryResponse
//   400: errorResponse
//   500: errorResponse

// swagger:route GET /repositories repository get-repos-routeid
//
// read all repositories
//
// responses:
//   200: stringArrayResponse
//   400: errorResponse
//   500: errorResponse

// swagger:response repositoryResponse
type repositoryResponseWrapper struct {
	// swagger:allOf
	models.Repository
}

// swagger:parameters add-repo-routeid
type repositoryCreationWrapper struct {
	// swagger:allOf
	Body models.RepositoryCreateRequest
}

// swagger:response stringArrayResponse
type stringArrayResponseWrapper struct {
	// in:body
	Body []string
}
