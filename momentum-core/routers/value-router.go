package routers

import (
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ROUTING_PATH_VALUE_BY_ID = VERSION + "/repository/:repositoryName/:valueId"
const ROUTING_PATH_VALUE = VERSION + "/value"

type ValueRouter struct {
	valueService *services.ValueService
}

func NewValueRouter(valueService *services.ValueService) *ValueRouter {
	return new(ValueRouter)
}

func (vr *ValueRouter) RegisterValueRoutes(server *gin.Engine) {

	server.GET(ROUTING_PATH_VALUE_BY_ID, vr.valueById)
}

func (vr *ValueRouter) valueById(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	repoName := c.Param("repositoryName")
	valueId := c.Param("valueId")

	result, err := vr.valueService.ValueById(repoName, valueId, traceId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, result)
}
