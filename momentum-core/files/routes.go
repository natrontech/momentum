package files

import (
	"momentum-core/artefacts"
	"momentum-core/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(engine *gin.Engine) {
	engine.GET(config.API_FILE_BY_ID, GetFile)
	engine.GET(config.API_DIR_BY_ID, GetDir)
	engine.GET(config.API_FILE_LINE_OVERWRITTENBY, GetOverwrittenBy)
	engine.GET(config.API_FILE_LINE_OVERWRITES, GetOverwrites)
}

// GetFile godoc
//
//	@Summary		gets the content of a file
//	@Tags			files
//	@Produce		json
//	@Param			id		path		string					true	"file id"
//	@Success		200		{object}	File
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/file/{id} [get]
func GetFile(c *gin.Context) {

	traceId := config.LOGGER.TraceId()
	id := c.Param("id")

	artefact, err := artefacts.FileById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, config.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	encodedFile, err := fileToBase64(artefacts.FullPath(artefact))
	file := NewFile(artefact.Id, artefact.Name, encodedFile)

	c.JSON(http.StatusOK, file)
}

// GetDir godoc
//
//	@Summary		gets the content of a file
//	@Tags			files
//	@Produce		json
//	@Param			id		path		string					true	"file id"
//	@Success		200		{object}	Dir
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/dir/{id} [get]
func GetDir(c *gin.Context) {

	_ = c.Param("id")

	return
}

// GetOverwrittenBy godoc
//
//	@Summary		gets a list of properties which overwrite the given line.
//	@Tags			files
//	@Produce		json
//	@Param			id				path		string	true	"file id"
//	@Param			lineNumber		path		int		true	"line number in file"
//	@Success		200		{array}		Overwrite
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/file/{id}/line/{lineNumber}/overwritten-by [get]
func GetOverwrittenBy(c *gin.Context) {

	_ = c.Param("id")
	_ = c.Param("lineNumber")
}

// GetOverwrites godoc
//
//	@Summary		gets a list of child properties, which are overwritten by the given line.
//	@Tags			files
//	@Produce		json
//	@Param			id				path		string	true	"file id"
//	@Param			lineNumber		path		int		true	"line number in file"
//	@Success		200		{array}		Overwrite
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/file/{id}/line/{lineNumber}/overwrites [get]
func GetOverwrites(c *gin.Context) {

	_ = c.Param("id")
	_ = c.Param("lineNumber")
}
