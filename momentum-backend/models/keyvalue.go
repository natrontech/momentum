package models

import (
	"momentum-core/clients"
	"momentum-core/tree"
	"momentum-core/utils"

	"github.com/gin-gonic/gin"
)

// swagger:model
type KeyValueCreateRequest struct {
	Key                string `json:"key"`
	Value              string `json:"value"`
	DisplayName        string `json:"displayName"`
	ParentStageId      string `json:"parentStageId"`
	ParentDeploymentId string `json:"parentDeploymentId"`
	ParentKeyValueId   string `json:"parentKeyValueId"`
}

// swagger:model
type KeyValue struct {
	Id                 string `json:"id"`
	Key                string `json:"key"`
	Value              string `json:"value"`
	Path               string `json:"-"`
	DisplayName        string `json:"displayName"`
	ParentStageId      string `json:"parentStageId"`
	ParentDeploymentId string `json:"parentDeploymentId"`
}

func ExtractKeyValueCreateRequest(c *gin.Context) (*KeyValueCreateRequest, error) {
	return utils.Extract[KeyValueCreateRequest](c)
}

func ExtractKeyValue(c *gin.Context) (*KeyValue, error) {
	return utils.Extract[KeyValue](c)
}

func ToKeyValueFromNode(n *tree.Node) (*KeyValue, error) {

	keyValue := new(KeyValue)

	keyValue.Id = n.Id
	keyValue.Key = n.FullPath()
	keyValue.Value = n.Value

	parentStage := n
	for !parentStage.IsStage() {
		if parentStage.Kind == tree.File {
			keyValue.ParentDeploymentId = parentStage.Id
		}
		parentStage = parentStage.Parent
	}
	keyValue.ParentStageId = parentStage.Id

	keyValue.DisplayName = n.Path

	return keyValue, nil
}

func ToKeyValue(data []byte) (*KeyValue, error) {

	keyValue, err := clients.UnmarshallJson[KeyValue](data)
	if err != nil {
		return nil, err
	}
	return keyValue, nil
}

func (kv *KeyValue) ToJson() ([]byte, error) {

	data, err := clients.MarshallJson(kv)
	if err != nil {
		return nil, err
	}
	return data, nil
}
