package templates

type TemplateStructureType int

const (
	DIR TemplateStructureType = 1 << iota
	FILE
)

type CreateTemplateRequest struct {
	Name              string                             `json:"name"`
	ParentId          string                             `json:"parentId"` // must be an artefact of type APPLICATION or STAGE
	TemplateStructure []*CreateTemplateStructureArtefact `json:"templateStructure"`
}

type CreateTemplateStructureArtefact struct {
	Name                  string                `json:"name"`
	Path                  string                `json:"path"` // the path is to be given relative to the templates root
	TemplateStructureType TemplateStructureType `json:"templateStructureType"`
	ParentId              string                `json:"parentId"` // if not set, it will live inside the templates root
}

type Template struct {
	Id               string              `json:"id"`
	Name             string              `json:"name"`
	TemplateArtefact []*TemplateArtefact `json:"templateArtefacts"`
}

type TemplateArtefact struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
