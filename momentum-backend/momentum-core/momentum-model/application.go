package momentummodel

import (
	"errors"

	"github.com/pocketbase/pocketbase/models"
)

const TABLE_APPLICATIONS_NAME = "applications"
const TABLE_APPLICATIONS_FIELD_ID = "id"
const TABLE_APPLICATIONS_FIELD_NAME = "name"
const TABLE_APPLICATIONS_FIELD_STAGES = "stages"
const TABLE_APPLICATIONS_FIELD_HELMREPO = "helmRepository"
const TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY = "parentRepository"

type IApplication interface {
	IModel

	Name() string
	SetName(string)

	ParentRepositoryId() string
	SetParentRepositoryId(string)

	StageIds() []string
	SetStageIds([]string)

	HelmRepositoryUrl() string
	SetHelmRepositoryUrl(string)
}

type Application struct {
	IApplication

	name               string
	parentRepositoryId string
	stageIds           []string
	helmRepositoryUrl  string
}

func ToApplication(record *models.Record) (IApplication, error) {

	if record.TableName() != TABLE_APPLICATIONS_NAME {
		return nil, errors.New("unallowed record type")
	}

	app := new(Application)
	app.SetId(record.Id)
	app.name = record.GetString(TABLE_APPLICATIONS_FIELD_NAME)
	app.stageIds = record.Get(TABLE_APPLICATIONS_FIELD_STAGES).([]string)
	app.helmRepositoryUrl = record.GetString(TABLE_APPLICATIONS_FIELD_HELMREPO)
	app.parentRepositoryId = record.GetString(TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY)

	return app, nil
}

func ToApplicationRecord(app IApplication, recordInstance *models.Record) (*models.Record, error) {

	if recordInstance.TableName() != TABLE_APPLICATIONS_NAME {
		return nil, errors.New("unallowed record type")
	}

	recordInstance.Set(TABLE_APPLICATIONS_FIELD_NAME, app.Name())
	recordInstance.Set(TABLE_APPLICATIONS_FIELD_STAGES, app.StageIds())
	recordInstance.Set(TABLE_APPLICATIONS_FIELD_HELMREPO, app.HelmRepositoryUrl())
	recordInstance.Set(TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY, app.ParentRepositoryId())

	return recordInstance, nil
}

func (a *Application) Name() string {
	return a.name
}

func (a *Application) SetName(name string) {
	a.name = name
}

func (a *Application) ParentRepositoryId() string {
	return a.parentRepositoryId
}

func (a *Application) SetParentRepositoryId(parentRepositoryId string) {
	a.parentRepositoryId = parentRepositoryId
}

func (a *Application) StageIds() []string {
	return a.stageIds
}

func (a *Application) SetStageIds(stageIds []string) {
	a.stageIds = stageIds
}

func (a *Application) HelmRepositoryUrl() string {
	return a.helmRepositoryUrl
}

func (a *Application) SetHelmRepositoryUrl(helmRepositoryUrl string) {
	a.helmRepositoryUrl = helmRepositoryUrl
}
