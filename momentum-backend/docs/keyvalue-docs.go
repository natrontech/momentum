package docs

import "momentum-core/models"

type keyValueResponseWrapper struct {
	// in:body
	Body models.KeyValue
}

type keyValueCreationWrapper struct {
	// in:body
	Body models.KeyValueCreateRequest
}
