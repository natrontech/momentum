package templates

import "momentum-core/files"

type TemplateStructureType int

type TemplateKind int

const (
	APPLICATION TemplateKind = 1 << iota
	STAGE
	DEPLOYMENT
)

const (
	DIR TemplateStructureType = 1 << iota
	FILE
)

type CreateTemplateRequest struct {
	TemplateKind TemplateKind `json:"templateKind"`
	Template     *files.Dir   `json:"template"`
}

type Template struct {
	TemplateKind TemplateKind `json:"templateKind"`
	Template     *files.Dir   `json:"template"`
}

type TemplateStore struct {
	Templates []*Template `json:"templates"`
}

// next steps:
//  1. define templates
//  2. define implement overwrite detection by file
//  3. implement backtracking with detected files for line matching endpoints

// this shall define which files overwrite each other.
// shall support wildcards and filenames.
type OverwriteConfiguration struct {
}
