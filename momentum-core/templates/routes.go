package templates

import (
	"momentum-core/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTemplateRoutes(engine *gin.Engine) {
	engine.GET(config.API_TEMPLATES_APPLICATIONS, GetApplicationTemplates)
	engine.GET(config.API_TEMPLATES_STAGES, GetStageTemplates)
	engine.GET(config.API_TEMPLATES_DEPLOYMENTS, GetDeploymentTemplates)
}

// GetApplicationTemplates godoc
//
//	@Summary		gets all available application templates
//	@Tags			templates
//	@Produce		json
//	@Success		200		{array}		Template
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
//	@Success		200		{array}		Template
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
//	@Success		200		{array}		Template
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates/deployments [get]
func GetDeploymentTemplates(c *gin.Context) {

	result := TemplateNames(config.DeploymentTemplatesPath(config.GLOBAL))
	c.JSON(http.StatusOK, result)
}

// GetTemplateSpec godoc
//
//	@Summary		gets the spec for a template, which contains values to be set when applying the template
//	@Tags			templates
//	@Produce		json
//	@Param			id		path		string			true	"file id"
//	@Success		200		{object}	Template
//	@Failure		400		{object}	config.ApiError
//	@Failure		404		{object}	config.ApiError
//	@Failure		500		{object}	config.ApiError
//	@Router			/api/beta/templates/spec/:templateName [get]
func GetTemplateSpec(c *gin.Context) {

}

func ApplyDeploymentTemplate(c *gin.Context) {

}
