package config

import "github.com/gin-gonic/gin"

// https://datatracker.ietf.org/doc/html/rfc7807
type ApiError struct {
	error

	ErrorType string `json:"type"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	Detail    string `json:"detail"`
	Instance  string `json:"instance"`
}

func NewApiError(err error, statusCode int, c *gin.Context, traceId string) *ApiError {

	e := new(ApiError)

	e.ErrorType = "api-error"
	e.Detail = traceId + " " + err.Error()
	e.Title = err.Error()
	e.Status = statusCode
	e.Instance = c.FullPath()

	return e
}
