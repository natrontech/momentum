package models

import (
	"errors"
	"momentum-core/clients"
	"momentum-core/tree"
	"momentum-core/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type DeploymentCreateRequest struct {
	Name              string `json:"name"`
	ReconcileInterval string `json:"reconcileInterval"`
	ChartVersion      string `json:"chartVersion"`
	ParentStageId     string `json:"parentStageId"`
	RepositoryName    string `json:"repositoryName"`
	ApplicationName   string `json:"applicationName"`
}

type Deployment struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Path          string `json:"-"`
	ParentStageId string `json:"parentStageId"`
	RepositoryId  string `json:"repositoryId"`
}

func ExtractDeploymentCreateRequest(c *gin.Context) (*DeploymentCreateRequest, error) {
	return utils.Extract[DeploymentCreateRequest](c)
}

func ExtractDeployment(c *gin.Context) (*Deployment, error) {
	return utils.Extract[Deployment](c)
}

func ToDeploymentFromNode(n *tree.Node, repositoryId string) (*Deployment, error) {

	if n == nil || n.Parent == nil {
		return nil, errors.New("nil node or parent")
	}

	deployment := new(Deployment)

	deployment.Id = n.Id
	deployment.Name = strings.Split(n.Path, "::")[0]
	deployment.Path = n.FullPath()
	deployment.RepositoryId = repositoryId
	deployment.ParentStageId = n.Parent.Id

	return deployment, nil
}

func ToDeplyoment(data []byte) (*Deployment, error) {

	deployment, err := clients.UnmarshallJson[Deployment](data)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

func (d *Deployment) ToJson() ([]byte, error) {

	data, err := clients.MarshallJson(d)
	if err != nil {
		return nil, err
	}
	return data, nil
}
