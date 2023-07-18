package models

import (
	"errors"
	"momentum-core/clients"
	"momentum-core/tree"
	"momentum-core/utils"

	"github.com/gin-gonic/gin"
)

type StageCreateRequest struct {
	Name                string `json:"name"`
	RepositoryName      string `json:"repositoryName"`
	ParentApplicationId string `json:"parentApplicationId"`
	ParentStageId       string `json:"parentStageId"`
}

type Stage struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Path                string `json:"-"`
	ParentApplicationId string `json:"parentApplicationId"`
	ParentStageId       string `json:"parentStageId"`
}

func ExtractStageCreateRequest(c *gin.Context) (*StageCreateRequest, error) {
	return utils.Extract[StageCreateRequest](c)
}

func ExtractStage(c *gin.Context) (*Stage, error) {
	return utils.Extract[Stage](c)
}

func ToStageFromNode(n *tree.Node, parentId string) (*Stage, error) {

	if n == nil {
		return nil, errors.New("nil node")
	}

	stage := new(Stage)
	stage.Id = n.Id
	stage.Name = n.Path
	stage.Path = n.FullPath()

	if n.IsStage() {
		stage.ParentStageId = parentId
	} else {
		stage.ParentApplicationId = parentId
	}

	return stage, nil
}

func ToStage(data []byte) (*Stage, error) {

	stage, err := clients.UnmarshallJson[Stage](data)
	if err != nil {
		return nil, err
	}
	return stage, nil
}

func (a *Stage) ToJson() ([]byte, error) {

	data, err := clients.MarshallJson(a)
	if err != nil {
		return nil, err
	}
	return data, nil
}
