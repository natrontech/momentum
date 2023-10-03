package templates

import (
	"momentum-core/files"
	"momentum-core/overwrites"
)

type TemplateKind int

const (
	APPLICATION TemplateKind = 1 << iota
	STAGE
	DEPLOYMENT
)

type TemplateDir struct {
	Name        string          `json:"name"`
	Directories []*TemplateDir  `json:"directories"`
	Files       []*TemplateFile `json:"files"`
}

type TemplateFile struct {
	Name         string `json:"name"`
	TemplateBody string `json:"templateBody"` // base64 encoded
}

type CreateTemplateRequest struct {
	TemplateKind TemplateKind `json:"templateKind"`
	// the toplevel directories name is the name of the template
	Template        *TemplateDir                `json:"template"`
	Children        *Template                   `json:"children"`
	OverwriteConfig *overwrites.OverwriteConfig `json:"overwriteConfig"`
}

type Template struct {
	// be aware, that each template must have an unique name
	// it doesn't matter if they are of different template kind
	TemplateName string       `json:"templateName"`
	TemplateKind TemplateKind `json:"templateKind"`
	Template     *files.Dir   `json:"template"`
	// The children are templates which are contained within the template.
	Children []*Template `json:"children"`
}

type TemplateSpec struct {
	ValueSpecs []*ValueSpec `json:"valueSpecs"`
}

type ValueSpec struct {
	TemplateId string `json:"templateId"` // id of the template which the value belongs to
	Name       string `json:"name"`       // name of the value (name displayed in frontend)
	Value      string `json:"value"`      // the value assigned
}
