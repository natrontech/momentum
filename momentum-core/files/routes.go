package files

import (
	"errors"
	"momentum-core/artefacts"
	"momentum-core/config"
	"momentum-core/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(engine *gin.Engine) {
	engine.GET(config.API_FILE_BY_ID, GetFile)
	engine.POST(config.API_FILE_ADD, AddFile)
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

// AddFile godoc
//
//	@Summary		adds a new file to a given parent
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Body 			CreateFileRequest
//	@Success		200		{object}	File
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/file [post]
func AddFile(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	createFileReq := new(CreateFileRequest)

	err := c.BindJSON(createFileReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	dir, err := artefacts.DirectoryById(createFileReq.ParentId)
	if err != nil {
		c.JSON(http.StatusNotFound, config.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	filePath := filepath.Join(artefacts.FullPath(dir), createFileReq.Name)
	if utils.FileExists(filePath) {
		err = errors.New("file with given name already exists")
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	fileContentDecoded, err := fileToRaw(createFileReq.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	writeSuccess := utils.FileWrite(filePath, fileContentDecoded)
	if !writeSuccess {
		err = errors.New("writing file failed")
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	encodedFile, err := fileToBase64(filePath)
	newFileId, err := utils.GenerateId(config.IdGenerationPath(filePath))
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}
	fileArtefact, err := artefacts.FileById(newFileId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	file := NewFile(fileArtefact.Id, fileArtefact.Name, encodedFile)

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
