package momentumservices

import (
	"errors"
	"fmt"
	consts "momentum/momentum-core/momentum-config"
	tree "momentum/momentum-core/momentum-tree"

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

func (as *ApplicationService) SyncApplicationsFromDisk(n *tree.Node, record *models.Record) ([]*models.Record, error) {

	recs := make([]*models.Record, 0)
	apps := n.Apps()
	for _, a := range apps {
		fmt.Println(a)
	}
	for _, app := range apps {

		stages, err := as.stageService.SyncStagesFromDisk(app)
		if err != nil {
			return nil, err
		}

		stageIds := make([]string, 0)
		for _, stage := range stages {
			stageIds = append(stageIds, stage.Id)
		}

		rec, err := as.createWithoutEvent(app.NormalizedPath(), stageIds)
		if err != nil {
			return nil, err
		}

		err = as.stageService.AddParentApplication(stageIds, rec)

		recs = append(recs, rec)
	}
	return recs, nil
}

func (as *ApplicationService) AddRepository(repositoryRecord *models.Record, applications []*models.Record) error {

	if repositoryRecord.Collection().Name != consts.TABLE_REPOSITORIES_NAME {
		return errors.New("repositoryRecord is not record of repositories collection")
	}

	for _, app := range applications {

		app.Set(consts.TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY, repositoryRecord.Id)
		err := as.saveWithoutEvent(app)
		if err != nil {
			return err
		}
	}

	return nil
}

func (as *ApplicationService) GetApplicationCollection() (*models.Collection, error) {

	return as.dao.FindCollectionByNameOrId(consts.TABLE_APPLICATIONS_NAME)
}

func (as *ApplicationService) createWithoutEvent(name string, stageIds []string) (*models.Record, error) {

	appCollection, err := as.GetApplicationCollection()
	if err != nil {
		return nil, err
	}

	appRecord := models.NewRecord(appCollection)
	appRecord.Set(consts.TABLE_APPLICATIONS_FIELD_NAME, name)
	appRecord.Set(consts.TABLE_APPLICATIONS_FIELD_STAGES, stageIds)

	return appRecord, as.saveWithoutEvent(appRecord)
}

func (as *ApplicationService) saveWithoutEvent(record *models.Record) error {
	return as.dao.Clone().SaveRecord(record)
}
