package momentummodel

import (
	"errors"

	"github.com/pocketbase/pocketbase/models"
)

const TABLE_DEPLOYMENTS_NAME = "deployments"
const TABLE_DEPLOYMENTS_FIELD_ID = "id"
const TABLE_DEPLOYMENTS_FIELD_NAME = "name"
const TABLE_DEPLOYMENTS_FIELD_DESCRIPTION = "description"
const TABLE_DEPLOYMENTS_FIELD_REPOSITORIES = "repositories"
const TABLE_DEPLOYMENTS_FIELD_KEYVALUES = "keyValues"
const TABLE_DEPLOYMENTS_FIELD_PARENTSTAGE = "parentStage"

type IDeployment interface {
	IModel

	Name() string
	SetName(string)

	Description() string
	SetDescription(string)

	ParentStageId() string
	SetParentStageId(string)

	KeyValueIds() []string
	SetKeyValueIds([]string)

	RepositoryIds() []string
	SetRepositoryIds([]string)
}

type Deployment struct {
	IDeployment

	name          string
	description   string
	parentStageId string
	keyValueIds   []string
	repositoryIds []string
}

func ToDeployment(record *models.Record) (IDeployment, error) {

	if record.TableName() != TABLE_DEPLOYMENTS_NAME {
		return nil, errors.New("unallowed record type")
	}

	dep := new(Deployment)
	dep.SetId(record.Id)
	dep.name = record.GetString(TABLE_APPLICATIONS_FIELD_NAME)
	dep.description = record.GetString(TABLE_DEPLOYMENTS_FIELD_DESCRIPTION)
	dep.parentStageId = record.GetString(TABLE_DEPLOYMENTS_FIELD_PARENTSTAGE)
	dep.keyValueIds = record.Get(TABLE_DEPLOYMENTS_FIELD_KEYVALUES).([]string)
	dep.repositoryIds = record.Get(TABLE_DEPLOYMENTS_FIELD_REPOSITORIES).([]string)

	return dep, nil
}

func ToDeploymentRecord(dep IDeployment, recordInstance *models.Record) (*models.Record, error) {

	if recordInstance.TableName() != TABLE_DEPLOYMENTS_NAME {
		return nil, errors.New("unallowed record type")
	}

	recordInstance.Set(TABLE_DEPLOYMENTS_FIELD_NAME, dep.Name())
	recordInstance.Set(TABLE_DEPLOYMENTS_FIELD_DESCRIPTION, dep.Description())
	recordInstance.Set(TABLE_DEPLOYMENTS_FIELD_REPOSITORIES, dep.RepositoryIds())
	recordInstance.Set(TABLE_DEPLOYMENTS_FIELD_KEYVALUES, dep.KeyValueIds())
	recordInstance.Set(TABLE_DEPLOYMENTS_FIELD_PARENTSTAGE, dep.ParentStageId())

	return recordInstance, nil
}

func (d *Deployment) Name() string {
	return d.name
}

func (d *Deployment) SetName(name string) {
	d.name = name
}

func (d *Deployment) Description() string {
	return d.description
}

func (d *Deployment) SetDescription(description string) {
	d.description = description
}

func (d *Deployment) ParentStageId() string {
	return d.parentStageId
}

func (d *Deployment) SetParentStageId(parentStageId string) {
	d.parentStageId = parentStageId
}

func (d *Deployment) KeyValueIds() []string {
	return d.keyValueIds
}

func (d *Deployment) SetKeyValueIds(keyValueIds []string) {
	d.keyValueIds = keyValueIds
}

func (d *Deployment) RepositoryIds() []string {
	return d.repositoryIds
}

func (d *Deployment) SetRepositoryIds(repositoryIds []string) {
	d.repositoryIds = repositoryIds
}
