package utils

import "github.com/gin-gonic/gin"

func Extract[T any](c *gin.Context) (*T, error) {

	t := new(T)
	err := c.BindJSON(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
