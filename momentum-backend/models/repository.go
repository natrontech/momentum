package models

import (
	"momentum-core/clients"
	"momentum-core/tree"
	"momentum-core/utils"

	"github.com/gin-gonic/gin"
)

// swagger:model
type RepositoryCreateRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// swagger:model
type Repository struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"-"`
}

func ExtractRepository(c *gin.Context) (*Repository, error) {
	return utils.Extract[Repository](c)
}

func ExtractRepositoryCreateRequest(c *gin.Context) (*RepositoryCreateRequest, error) {
	return utils.Extract[RepositoryCreateRequest](c)
}

func ToRepositoryFromNode(n *tree.Node) (*Repository, error) {

	repository := new(Repository)

	repository.Id = n.Id
	repository.Name = n.Path
	repository.Path = n.FullPath()

	return repository, nil
}

func ToRepository(data []byte) (*Repository, error) {

	repository, err := clients.UnmarshallJson[Repository](data)
	if err != nil {
		return nil, err
	}
	return repository, nil
}

func (r *Repository) ToJson() ([]byte, error) {

	data, err := clients.MarshallJson(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}
