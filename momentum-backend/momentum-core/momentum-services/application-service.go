package momentumservices

import (
	"errors"
	model "momentum/momentum-core/momentum-model"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type ApplicationService struct {
	dao          *daos.Dao
	stageService *StageService
}

func NewApplicationService(dao *daos.Dao, stageService *StageService) *ApplicationService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	appService := new(ApplicationService)
	appService.dao = dao
	appService.stageService = stageService

	return appService
}

func (as *ApplicationService) AddRepository(repositoryRecord *models.Record, applications []*models.Record) error {

	if repositoryRecord.Collection().Name != model.TABLE_REPOSITORIES_NAME {
		return errors.New("repositoryRecord is not record of repositories collection")
	}

	for _, app := range applications {

		app.Set(model.TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY, repositoryRecord.Id)
		err := as.saveWithoutEvent(app)
		if err != nil {
			return err
		}
	}

	return nil
}

func (as *ApplicationService) GetApplicationCollection() (*models.Collection, error) {

	return as.dao.FindCollectionByNameOrId(model.TABLE_APPLICATIONS_NAME)
}

func (as *ApplicationService) createWithoutEvent(name string, stageIds []string) (*models.Record, error) {

	appCollection, err := as.GetApplicationCollection()
	if err != nil {
		return nil, err
	}

	appRecord := models.NewRecord(appCollection)
	appRecord.Set(model.TABLE_APPLICATIONS_FIELD_NAME, name)
	appRecord.Set(model.TABLE_APPLICATIONS_FIELD_STAGES, stageIds)

	return appRecord, as.saveWithoutEvent(appRecord)
}

func (as *ApplicationService) saveWithoutEvent(record *models.Record) error {
	return as.dao.Clone().SaveRecord(record)
}
