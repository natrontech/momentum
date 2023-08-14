package models

import (
	"errors"
	"momentum-core/clients"
	"momentum-core/tree"
	"momentum-core/utils"

	"github.com/gin-gonic/gin"
)

type ValueType int

const (
	REPOSITORY ValueType = iota
	APPLICATION
	STAGE
	DEPLOYMENT
)

type KeyValueCreateRequest struct {
	Key                string `json:"key"`
	Value              string `json:"value"`
	DisplayName        string `json:"displayName"`
	ParentStageId      string `json:"parentStageId"`
	ParentDeploymentId string `json:"parentDeploymentId"`
	ParentKeyValueId   string `json:"parentKeyValueId"`
}

type Value struct {
	Id                 string `json:"id"`
	Key                string `json:"key"`
	Value              string `json:"value"`
	Path               string `json:"-"`
	DisplayName        string `json:"displayName"`
	ParentStageId      string `json:"parentStageId"`
	ParentDeploymentId string `json:"parentDeploymentId"`
}

type ValueWrapper struct {
	FileId    string    `json:"parentFileId"`
	FileName  string    `json:"parentFileName"`
	ValueType ValueType `json:"valueType"`
	Values    []*Value  `json:"values"`
}

func ExtractKeyValueCreateRequest(c *gin.Context) (*KeyValueCreateRequest, error) {
	return utils.Extract[KeyValueCreateRequest](c)
}

func ExtractKeyValue(c *gin.Context) (*Value, error) {
	return utils.Extract[Value](c)
}

func ToValueWrapperFromNode(n *tree.Node, valType ValueType) (*ValueWrapper, error) {

	if n.Kind != tree.File {
		return nil, errors.New("only files can be converted to a value wrapper")
	}

	valueWrapper := new(ValueWrapper)

	valueWrapper.FileId = n.Id
	valueWrapper.FileName = n.NormalizedPath()
	mappedValues := make([]*Value, 0)
	for _, v := range n.Values() {
		value, err := ToValueFromNode(v)
		if err != nil {
			return nil, errors.New("unable to map value from node")
		}
		mappedValues = append(mappedValues, value)
	}
	valueWrapper.Values = mappedValues
	valueWrapper.ValueType = valType

	return valueWrapper, nil
}

func ToValueFromNode(n *tree.Node) (*Value, error) {

	keyValue := new(Value)

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

func ToKeyValue(data []byte) (*Value, error) {

	keyValue, err := clients.UnmarshallJson[Value](data)
	if err != nil {
		return nil, err
	}
	return keyValue, nil
}

func (kv *Value) ToJson() ([]byte, error) {

	data, err := clients.MarshallJson(kv)
	if err != nil {
		return nil, err
	}
	return data, nil
}
