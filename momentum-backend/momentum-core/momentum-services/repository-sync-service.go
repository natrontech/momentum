package momentumservices

import (
	config "momentum/momentum-core/momentum-config"
	tree "momentum/momentum-core/momentum-tree"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type RepositorySyncService struct {
	dao                *daos.Dao
	applicationService *ApplicationService
}

func (rs *RepositorySyncService) SyncRepositoryFromDisk(n *tree.Node, record *models.Record) (*models.Record, []*models.Record, []*models.Record, error) {

	appRecords, err := rs.applicationService.SyncApplicationsFromDisk(n, record)
	if err != nil {
		return nil, nil, nil, apis.NewApiError(500, err.Error(), nil)
	}

	appRecIds := make([]string, 0)
	for _, appRec := range appRecords {
		appRecIds = append(appRecIds, appRec.Id)
	}
	record.Set(config.TABLE_REPOSITORIES_FIELD_APPLICATIONS, appRecIds)

	// this complex loop is necessary because we need to know which deployments must add the repository
	// which is currently created, when the creation of the repository is finished.
	// TODO for a future refactoring: extract logic to specific services.
	deployments := make([]*models.Record, 0)
	for _, applicationRecord := range appRecords {

		appRecord, err := rs.dao.FindRecordById(config.TABLE_APPLICATIONS_NAME, applicationRecord.Id)
		if err != nil {
			return nil, nil, nil, err
		}

		stagesIds := appRecord.Get(config.TABLE_APPLICATIONS_FIELD_STAGES).([]string)
		for _, stageId := range stagesIds {

			stageRec, err := rs.dao.FindRecordById(config.TABLE_STAGES_NAME, stageId)
			if err != nil {
				return nil, nil, nil, err
			}

			deploymentIds := stageRec.Get(config.TABLE_STAGES_FIELD_DEPLOYMENTS).([]string)
			for _, deploymentId := range deploymentIds {

				deploymentRec, err := rs.dao.FindRecordById(config.TABLE_DEPLOYMENTS_NAME, deploymentId)
				if err != nil {
					return nil, nil, nil, err
				}

				deployments = append(deployments, deploymentRec)
			}
		}
	}

	return record, appRecords, deployments, nil
}
