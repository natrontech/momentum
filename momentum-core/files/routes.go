package files

import (
	"errors"
	"momentum-core/artefacts"
	"momentum-core/backtracking"
	"momentum-core/config"
	"momentum-core/utils"
	"momentum-core/yaml"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	gittransaction "github.com/Joel-Haeberli/git-transaction"
	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(engine *gin.Engine) {
	engine.GET(config.API_FILE_BY_ID, GetFile)
	engine.POST(config.API_FILE_ADD, AddFile)
	engine.GET(config.API_FILE_LINE_OVERWRITTENBY, GetOverwrittenBy)
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
//	@Summary		adds a new file to a given parent (triggers transaction)
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

	ctx, transaction, err := gittransaction.New(config.TRANSACTION_MODE, config.GLOBAL.RepoDir(), config.GLOBAL.TransactionToken())
	if err != nil {
		transaction.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	writeSuccess := utils.FileWrite(filePath, fileContentDecoded)
	if !writeSuccess {
		transaction.Rollback(ctx)
		err = errors.New("writing file failed")
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Write(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	encodedFile, err := fileToBase64(filePath)
	newFileId, err := utils.GenerateId(config.IdGenerationPath(filePath))
	if err != nil {
		transaction.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	fileArtefact, err := artefacts.FileById(newFileId)
	if err != nil {
		transaction.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	file := NewFile(fileArtefact.Id, fileArtefact.Name, encodedFile)

	c.JSON(http.StatusOK, file)
}

// UpdateFile godoc
//
//	@Summary		updates the given file (triggers transaction)
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Body 			File
//	@Success		200		{object}	File
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/file/{id} [put]
func UpdateFile(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	requestedFile := new(File)
	err := c.BindJSON(requestedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	decodedBody, err := fileToRaw(requestedFile.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	file, err := artefacts.FileById(requestedFile.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, config.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	filePath := artefacts.FullPath(file)
	if !utils.FileExists(filePath) {
		err = errors.New("file does not exist")
		c.JSON(http.StatusNotFound, config.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	ctx, transaction, err := gittransaction.New(config.TRANSACTION_MODE, config.GLOBAL.RepoDir(), config.GLOBAL.TransactionToken())
	if err != nil {
		transaction.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	utils.FileDelete(filePath)
	success := utils.FileWrite(filePath, decodedBody)
	if !success {
		err = errors.New("failed to write file")
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Write(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	err = transaction.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	c.JSON(http.StatusOK, requestedFile)
}

// GetOverwrittenBy godoc
//
//	@Summary		gets a list of overwrites which overwrite the given line.
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

	traceId := config.LOGGER.TraceId()

	overwritableId := c.Param("id")
	overwritableLine, err := strconv.Atoi(c.Param("lineNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	overwritable, err := artefacts.FileById(overwritableId)
	if err != nil {
		c.JSON(http.StatusNotFound, config.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	fileNode, err := yaml.ParseFile(artefacts.FullPath(overwritable))
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	lineNode := yaml.FindNodeByLine(fileNode, overwritableLine)
	if lineNode == nil {
		err := errors.New("could not find line " + strconv.Itoa(overwritableLine) + " in file " + artefacts.FullPath(overwritable))
		c.JSON(http.StatusNotFound, config.NewApiError(err, http.StatusNotFound, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	overwritingFiles := make([]*artefacts.Artefact, 0)
	for _, advice := range artefacts.ActiveOverwriteAdvices {
		overwritingFiles = append(overwritingFiles, advice(artefacts.FullPath(overwritable))...)
	}

	if len(overwritingFiles) > 0 {

		predicate := yaml.ToMatchableSearchTerm(lineNode.FullPath())
		predicate = strings.Join(strings.Split(predicate, ".")[1:], ".") // remove filename prefix

		overwrites := make([]*Overwrite, 0)
		for _, overwriting := range overwritingFiles {

			backtracker := backtracking.NewPropertyBacktracker(predicate, artefacts.FullPath(overwriting), backtracking.NewYamlPropertyParser())
			var result []*backtracking.Match[string, yaml.ViewNode] = backtracker.RunBacktrack()

			for _, match := range result {

				overwrite := new(Overwrite)
				overwrite.OriginFileId = overwritableId
				overwrite.OriginFileLine = overwritableLine
				overwrite.OverwriteFileLine = match.MatchNode.Pointer.YamlNode.Line
				overwrite.OverwriteFileId = overwriting.Id

				overwrites = append(overwrites, overwrite)
			}
		}

		c.JSON(http.StatusOK, overwrites)
		return
	}

	c.JSON(http.StatusOK, make([]*Overwrite, 0))
}
