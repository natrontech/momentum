package docs

import "momentum-core/models"

type stageResponseWrapper struct {
	// in:body
	Body models.Stage
}

type stageCreationWrapper struct {
	// in:body
	Body models.StageCreateRequest
}
