package artefacts

import "momentum-core/tree"

type ArtefactType int

const (
	ROOT ArtefactType = 1 << iota
	META
	APPLICATION
	STAGE
	DEPLOYMENT
	FILE
)

type Artefact struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	ArtefactType ArtefactType `json:"type"`
	ContentIds   []string     `json:"contentIds"` // id's of children artefacts
	Content      []*Artefact  `json:"-"`
	ParentId     string       `json:"parentId"` // id of parent artefacts
	Parent       *Artefact    `json:"-"`
}

func toArtefact(n *tree.Node) *Artefact {

	if n == nil {
		return nil
	}

	artefact := new(Artefact)
	artefact.Id = n.Id
	artefact.Name = n.NormalizedPath()

	if n.Kind == tree.File {
		artefact.ArtefactType = FILE
	}

	return artefact
}
