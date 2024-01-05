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

type TemplateConfig struct {
	Kind     TemplateKind `json:"kind" yaml:"kind"`
	Children []string     `json:"children" yaml:"children"`
}

type CreateTemplateRequest struct {
	TemplateKind TemplateKind `json:"templateKind"`
	// the toplevel directories name is the name of the template
	Template        *TemplateDir                `json:"template"`
	TemplateConfig  *TemplateConfig             `json:"templateConfig"`
	OverwriteConfig *overwrites.OverwriteConfig `json:"-"` // `json:"overwriteConfig"` // TODO: implemented later
}

type Template struct {
	// be aware, that each template must have an unique name
	// it doesn't matter if they are of different template kind
	Name string       `json:"name"`
	Kind TemplateKind `json:"kind"`
	Root *files.Dir   `json:"root"`
	// The children are templates which are contained within the template.
	Children []*Template `json:"children"`
}

type TemplateSpec struct {
	Template   *Template    `json:"template"`
	ValueSpecs []*ValueSpec `json:"valueSpecs"`
}

type ValueSpec struct {
	TemplateName string `json:"templateName"` // name of the template which the value belongs to
	Name         string `json:"name"`         // name of the value (name displayed in frontend)
	Value        string `json:"value"`        // the value assigned
}
