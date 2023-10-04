package templates

import (
	"errors"
	"momentum-core/config"
	"net/http"

	gittransaction "github.com/Joel-Haeberli/git-transaction"
	"github.com/gin-gonic/gin"
)

func RegisterTemplateRoutes(engine *gin.Engine) {
	engine.GET(config.API_TEMPLATES_APPLICATIONS, GetApplicationTemplates)
	engine.GET(config.API_TEMPLATES_STAGES, GetStageTemplates)
	engine.GET(config.API_TEMPLATES_DEPLOYMENTS, GetDeploymentTemplates)
	engine.POST(config.API_TEMPLATES_ADD, AddTemplate)
	engine.GET(config.API_TEMPLATE_GET_SPEC, GetTemplateSpec)
	engine.POST(config.API_TEMPLATE_APPLY, ApplyDeploymentTemplate)
}

// GetApplicationTemplates godoc
//
//	@Summary		gets all available application templates
//	@Tags			templates
//	@Produce		json
//	@Success		200		{array}		string
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates/applications [get]
func GetApplicationTemplates(c *gin.Context) {

	result := TemplateNames(config.ApplicationTemplatesPath(config.GLOBAL))
	c.JSON(http.StatusOK, result)
}

// GetStageTemplates godoc
//
//	@Summary		gets all available stage templates
//	@Tags			templates
//	@Produce		json
//	@Success		200		{array}		string
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates/stages [get]
func GetStageTemplates(c *gin.Context) {

	result := TemplateNames(config.StageTemplatesPath(config.GLOBAL))
	c.JSON(http.StatusOK, result)
}

// GetDeploymentTemplates godoc
//
//	@Summary		gets all available deployment templates
//	@Tags			templates
//	@Produce		json
//	@Success		200		{array}		string
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates/deployments [get]
func GetDeploymentTemplates(c *gin.Context) {

	result := TemplateNames(config.DeploymentTemplatesPath(config.GLOBAL))
	c.JSON(http.StatusOK, result)
}

// AddTemplate godoc
//
//	@Summary		adds a new template (triggers transaction)
//	@Tags			templates
//	@Accept			json
//	@Produce		json
//	@Body			json
//	@Param 			CreateTemplateRequest body CreateTemplateRequest true "the body shall contain a CreateTemplateRequest instance"
//	@Success		200		{object}	Template
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates [post]
func AddTemplate(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	request := new(CreateTemplateRequest)
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	if TemplateExists(request.Template.Name) {
		err = errors.New("template with this name already exists")
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	if len(request.Template.Directories) == 0 && len(request.Template.Files) == 0 {
		err = errors.New("expecting non empty template")
		c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	validConfig, err := validTemplateConfig(request.TemplateConfig, request.TemplateKind)
	if !validConfig {
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

	temp, err := CreateTemplate(request)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, config.NewApiError(err, http.StatusBadRequest, c, traceId))
			config.LOGGER.LogError(err.Error(), err, traceId)
			return
		}
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

	c.JSON(http.StatusOK, temp)
}

// GetTemplateSpec godoc
//
//	@Summary		gets the spec for a template, which contains values to be set when applying the template
//	@Tags			templates
//	@Produce		json
//	@Param			templateName		path		string			true	"name of the template (template names are unique)"
//	@Success		200		{object}	Template
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates/spec/:templateName [get]
func GetTemplateSpec(c *gin.Context) {

	c.JSON(http.StatusNoContent, config.NewApiError(errors.New("not implemented yet"), http.StatusNoContent, c, config.LOGGER.TraceId()))
}

// ApplyDeploymentTemplate godoc
//
//	@Summary		gets the spec for a template, which contains values to be set when applying the template (triggers transaction)
//	@Tags			templates
//	@Accept 		json
//	@Produce		json
//	@Body			json
//	@Param			anchorArtefactId		path		string					true	"id of the artefact where the template shall be applied. Must be a directory."
//	@Param 			TemplateSpec 			body 		TemplateSpec 			true 	"the body shall contain a CreateTemplateRequest instance"
//	@Success		200		{object}	Template
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates/spec/apply/:anchorArtefactId [post]
func ApplyDeploymentTemplate(c *gin.Context) {

	traceId := config.LOGGER.TraceId()

	ctx, transaction, err := gittransaction.New(config.TRANSACTION_MODE, config.GLOBAL.RepoDir(), config.GLOBAL.TransactionToken())
	if err != nil {
		transaction.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, config.NewApiError(err, http.StatusInternalServerError, c, traceId))
		config.LOGGER.LogError(err.Error(), err, traceId)
		return
	}

	// TODO APPLY TEMPLATE HERE

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

	c.JSON(http.StatusNoContent, config.NewApiError(errors.New("not implemented yet"), http.StatusNoContent, c, config.LOGGER.TraceId()))
}
