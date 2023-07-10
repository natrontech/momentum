package momentummodel

import (
	"errors"

	"github.com/pocketbase/pocketbase/models"
)

const TABLE_REPOSITORIES_NAME = "repositories"
const TABLE_REPOSITORIES_FIELD_ID = "id"
const TABLE_REPOSITORIES_FIELD_NAME = "name"
const TABLE_REPOSITORIES_FIELD_URL = "url"
const TABLE_REPOSITORIES_FIELD_APPLICATIONS = "applications"

type IRepository interface {
	IModel

	Name() string
	SetName(string)

	Url() string
	SetUrl(string)

	ApplicationIds() []string
	SetApplicationIds([]string)
}

type Repository struct {
	IRepository
	Model

	name           string
	url            string
	applicationIds []string
}

func ToRepository(record *models.Record) (IRepository, error) {

	if record.TableName() != TABLE_REPOSITORIES_NAME {
		return nil, errors.New("unallowed record type")
	}

	repo := new(Repository)
	repo.SetId(record.Id)
	repo.name = record.GetString(TABLE_REPOSITORIES_FIELD_NAME)
	repo.url = record.GetString(TABLE_REPOSITORIES_FIELD_URL)
	repo.applicationIds = record.Get(TABLE_REPOSITORIES_FIELD_APPLICATIONS).([]string)

	return repo, nil
}

func ToRepositoryRecord(repo IRepository, recordInstance *models.Record) (*models.Record, error) {

	if recordInstance.TableName() != TABLE_REPOSITORIES_NAME {
		return nil, errors.New("unallowed record type")
	}

	recordInstance.Set(TABLE_REPOSITORIES_FIELD_NAME, repo.Name())
	recordInstance.Set(TABLE_REPOSITORIES_FIELD_URL, repo.Url())
	recordInstance.Set(TABLE_REPOSITORIES_FIELD_APPLICATIONS, repo.ApplicationIds())

	return recordInstance, nil
}

func (r *Repository) Id() string {
	return r.id
}

func (r *Repository) SetId(id string) {
	r.id = id
}

func (r *Repository) Name() string {
	return r.name
}

func (r *Repository) SetName(name string) {
	r.name = name
}

func (r *Repository) Url() string {
	return r.url
}

func (r *Repository) SetUrl(url string) {
	r.url = url
}

func (r *Repository) ApplicationIds() []string {
	return r.applicationIds
}

func (r *Repository) SetApplicationIds(applicationIds []string) {
	r.applicationIds = applicationIds
}
