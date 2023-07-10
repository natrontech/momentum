package momentummodel

import (
	"errors"

	"github.com/pocketbase/pocketbase/models"
)

const TABLE_KEYVALUE_NAME = "keyValues"
const TABLE_KEYVALUE_FIELD_ID = "id"
const TABLE_KEYVALUE_FIELD_KEY = "key"
const TABLE_KEYVALUE_FIELD_VALUE = "value"
const TABLE_KEYVALUE_FIELD_DISPLAY_NAME = "displayName"
const TABLE_KEYVALUE_FIELD_PARENTSTAGE = "parentStage"
const TABLE_KEYVALUE_FIELD_PARENTDEPLOYMENT = "parentDeployment"

type IKeyValue interface {
	IModel

	Key() string
	SetKey(string)

	Value() string
	SetValue(string)

	ParentStageId() string
	SetParentStageId(string)

	ParentDeploymentId() string
	SetParentDeploymentId(string)

	DisplayName() string
	SetDisplayName(string)
}

type KeyValue struct {
	IKeyValue
	Model

	key                string
	value              string
	parentStageId      string
	parentDeploymentId string
	displayName        string
}

func ToKeyValue(record *models.Record) (IKeyValue, error) {

	if record.TableName() != TABLE_KEYVALUE_NAME {
		return nil, errors.New("unallowed record type")
	}

	kv := new(KeyValue)
	if record.Id != "" {
		kv.SetId(record.Id)
	}
	kv.key = record.GetString(TABLE_KEYVALUE_FIELD_KEY)
	kv.value = record.GetString(TABLE_KEYVALUE_FIELD_VALUE)
	kv.displayName = record.GetString(TABLE_KEYVALUE_FIELD_DISPLAY_NAME)
	kv.parentStageId = record.GetString(TABLE_KEYVALUE_FIELD_PARENTSTAGE)
	kv.parentDeploymentId = record.GetString(TABLE_KEYVALUE_FIELD_PARENTDEPLOYMENT)

	return kv, nil
}

func ToKeyValueRecord(kv IKeyValue, recordInstance *models.Record) (*models.Record, error) {

	if recordInstance.TableName() != TABLE_KEYVALUE_NAME {
		return nil, errors.New("unallowed record type")
	}

	recordInstance.Set(TABLE_KEYVALUE_FIELD_KEY, kv.Key())
	recordInstance.Set(TABLE_KEYVALUE_FIELD_VALUE, kv.Value())
	recordInstance.Set(TABLE_KEYVALUE_FIELD_DISPLAY_NAME, kv.DisplayName())
	recordInstance.Set(TABLE_KEYVALUE_FIELD_PARENTSTAGE, kv.ParentStageId())
	recordInstance.Set(TABLE_KEYVALUE_FIELD_PARENTDEPLOYMENT, kv.ParentDeploymentId())

	return recordInstance, nil
}

func (kv *KeyValue) Id() string {
	return kv.id
}

func (kv *KeyValue) SetId(id string) {
	kv.id = id
}

func (kv *KeyValue) Key() string {
	return kv.key
}

func (kv *KeyValue) SetKey(key string) {
	kv.key = key
}

func (kv *KeyValue) Value() string {
	return kv.value
}

func (kv *KeyValue) SetValue(value string) {
	kv.value = value
}

func (kv *KeyValue) ParentStageId() string {
	return kv.parentStageId
}

func (kv *KeyValue) SetParentStageId(parentStageId string) {
	kv.parentStageId = parentStageId
}

func (kv *KeyValue) ParentDeploymentId() string {
	return kv.parentDeploymentId
}

func (kv *KeyValue) SetParentDeploymentId(parentDeploymentId string) {
	kv.parentDeploymentId = parentDeploymentId
}

func (kv *KeyValue) DisplayName() string {
	return kv.displayName
}

func (kv *KeyValue) SetDisplayName(displayName string) {
	kv.displayName = displayName
}
