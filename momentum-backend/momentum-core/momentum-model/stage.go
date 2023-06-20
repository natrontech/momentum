package momentummodel

import (
	"errors"

	"github.com/pocketbase/pocketbase/models"
)

const TABLE_STAGES_NAME = "stages"
const TABLE_STAGES_FIELD_ID = "id"
const TABLE_STAGES_FIELD_NAME = "name"
const TABLE_STAGES_FIELD_DEPLOYMENTS = "deployments"
const TABLE_STAGES_FIELD_KEYVALUES = "keyValues"
const TABLE_STAGES_FIELD_PARENTSTAGE = "parentStage"
const TABLE_STAGES_FIELD_PARENTAPPLICATION = "parentApplication"

type IStage interface {
	IModel

	Name() string
	SetName(string)

	ParentApplicationId() string
	SetParentApplicationId(string)

	ParentStageId() string
	SetParentStageId(string)

	DeploymentIds() []string
	SetDeploymentIds([]string)

	KeyValueIds() []string
	SetKeyValueIds([]string)
}

type Stage struct {
	IStage

	name                string
	parentApplicationId string
	parentStageId       string
	deploymentIds       []string
	keyValueIds         []string
}

func ToStage(record *models.Record) (IStage, error) {

	if record.TableName() != TABLE_STAGES_NAME {
		return nil, errors.New("unallowed record type")
	}

	st := new(Stage)
	st.SetId(record.Id)
	st.name = record.GetString(TABLE_STAGES_FIELD_NAME)
	st.parentApplicationId = record.GetString(TABLE_STAGES_FIELD_PARENTAPPLICATION)
	st.parentStageId = record.GetString(TABLE_STAGES_FIELD_PARENTSTAGE)
	st.deploymentIds = record.Get(TABLE_STAGES_FIELD_DEPLOYMENTS).([]string)
	st.keyValueIds = record.Get(TABLE_STAGES_FIELD_KEYVALUES).([]string)

	return st, nil
}

func ToStageRecord(st IStage, recordInstance *models.Record) (*models.Record, error) {

	if recordInstance.TableName() != TABLE_STAGES_NAME {
		return nil, errors.New("unallowed record type")
	}

	recordInstance.Set(TABLE_STAGES_FIELD_NAME, st.Name())
	recordInstance.Set(TABLE_STAGES_FIELD_PARENTSTAGE, st.ParentStageId())
	recordInstance.Set(TABLE_STAGES_FIELD_PARENTAPPLICATION, st.ParentApplicationId())
	recordInstance.Set(TABLE_STAGES_FIELD_KEYVALUES, st.KeyValueIds())
	recordInstance.Set(TABLE_STAGES_FIELD_DEPLOYMENTS, st.DeploymentIds())

	return recordInstance, nil
}

func (s *Stage) Name() string {
	return s.name
}

func (s *Stage) SetName(name string) {
	s.name = name
}

func (s *Stage) ParentApplicationId() string {
	return s.parentApplicationId
}

func (s *Stage) SetParentApplicationId(parentApplicationId string) {
	s.parentApplicationId = parentApplicationId
}

func (s *Stage) ParentStageId() string {
	return s.parentStageId
}

func (s *Stage) SetParentStageId(parentStageId string) {
	s.parentStageId = parentStageId
}

func (s *Stage) DeploymentIds() []string {
	return s.deploymentIds
}

func (s *Stage) SetDeploymentIds(deploymentIds []string) {
	s.deploymentIds = deploymentIds
}

func (s *Stage) KeyValueIds() []string {
	return s.keyValueIds
}

func (s *Stage) SetKeyValueIds(keyValueIds []string) {
	s.keyValueIds = keyValueIds
}
