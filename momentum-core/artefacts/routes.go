package artefacts

import (
	"errors"
	"momentum-core/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterArtefactRoutes(engine *gin.Engine) {
	engine.GET(config.API_ARTEFACT_BY_ID, GetArtefact)
	engine.GET(config.API_APPLICATIONS, GetApplications)
	engine.GET(config.API_STAGES, GetStages)
	engine.GET(config.API_DEPLOYMENTS, GetDeployments)
}

// GetArtefact godoc
//
//	@Summary		get an artefact by id (an artefact can be any of APPLICATION, STAGE or DEPLOYMENT)
//	@Tags			artefacts
//	@Produce		json
//	@Param			id		path		string					true	"artefact id"
//	@Success		200		{object}	Artefact
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/artefact/{id}/ [get]
func GetArtefact(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	artefactId := c.Param("id")
	if artefactId == "" {
		c.JSON(http.StatusNotFound, config.NewApiError(errors.New("artefact id not specified"), http.StatusBadRequest, c, traceId))
		return
	}

	t, err := LoadArtefactTree()
	if err != nil {
		c.JSON(http.StatusNotFound, config.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	// fmt.Println(WriteToString(t))

	flattened := FlatPreorder(t, make([]*Artefact, 0))
	for _, artefact := range flattened {
		if artefact.Id == artefactId {
			// fmt.Println(artefact, "artefactId:", artefact.ParentId)
			c.JSON(http.StatusOK, artefact)
			return
		}
	}

	return
}

// GetApplications godoc
//
//	@Summary		gets a list of all applications
//	@Tags			artefacts
//	@Produce		json
//	@Success		200		{array}		Artefact
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/applications [get]
func GetApplications(c *gin.Context) {

	c.JSON(http.StatusOK, Applications())
}

// GetStages godoc
//
//	@Summary		gets a list of all stages within an application or stage by id.
//	@Tags			artefacts
//	@Produce		json
//	@Param			parentId		query		string					true	"application or stage id"
//	@Success		200		{object}	Artefact
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/stages [get]
func GetStages(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	appOrStageId := c.Query("parentId")

	if appOrStageId == "" {
		err := errors.New("request parameter parentId not set")
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, StagesByApplication(appOrStageId))
}

// GetDeployments godoc
//
//	@Summary		get a list of deployments for a given stage by id
//	@Tags			artefacts
//	@Produce		json
//	@Param			stageId				query		string					true	"stage id"
//	@Success		200		{object}	Artefact
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/deployments [get]
func GetDeployments(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	stageId := c.Query("stageId")

	if stageId == "" {
		err := errors.New("request parameter stageId not set")
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, DeploymentsByStage(stageId))
}
