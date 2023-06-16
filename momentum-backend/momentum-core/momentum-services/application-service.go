package momentumservices

import (
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

func (as *ApplicationService) SyncApplicationsFromDisk(n *tree.Node, record *models.Record) ([]string, error) {

	recs := make([]string, 0)
	apps := n.Apps()
	for _, app := range apps {

		stages, err := as.stageService.SyncStagesFromDisk(app)
		if err != nil {
			return nil, err
		}

		rec, err := as.createWithoutEvent(app.NormalizedPath(), stages)
		if err != nil {
			return nil, err
		}

		err = as.stageService.AddParentApplication(stages, rec)

		recs = append(recs, rec.Id)
	}
	return recs, nil
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

	err = as.dao.Clone().SaveRecord(appRecord)

	return appRecord, err
}
