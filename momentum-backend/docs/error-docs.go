package docs

import "momentum-core/models"

// swagger:response errorResponse
type errorWrapper struct {
	// in:body
	Body models.ApiError
}
