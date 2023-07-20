package models

import (
	"errors"
	"momentum-core/clients"
	"momentum-core/tree"
	"momentum-core/utils"

	"github.com/gin-gonic/gin"
)

type ApplicationCreateRequest struct {
	Name              string `json:"name"`
	ReconcileInterval string `json:"reconcileInterval"`
	ChartVersion      string `json:"chartVersion"`
	RepositoryName    string `json:"repositoryName"`
}

type Application struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"-"`
	RepositoryName string `json:"repositoryName"`
}

func ExtractApplicationCreateRequest(c *gin.Context) (*ApplicationCreateRequest, error) {
	return utils.Extract[ApplicationCreateRequest](c)
}

func ExtractApplication(c *gin.Context) (*Application, error) {
	return utils.Extract[Application](c)
}

func ToApplicationFromNode(n *tree.Node, repositoryName string) (*Application, error) {

	if n == nil {
		return nil, errors.New("nil node")
	}

	app := new(Application)

	app.Id = n.Id
	app.Name = n.Path
	app.Path = n.FullPath()
	app.RepositoryName = repositoryName

	return app, nil
}

func ToApplication(data []byte) (*Application, error) {

	app, err := clients.UnmarshallJson[Application](data)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *Application) ToJson() ([]byte, error) {

	data, err := clients.MarshallJson(a)
	if err != nil {
		return nil, err
	}
	return data, nil
}
