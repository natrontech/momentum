package momentumservices

import (
	"fmt"
	consts "momentum/momentum-core/momentum-config"
	tree "momentum/momentum-core/momentum-tree"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type RepositoryService struct {
	dao                *daos.Dao
	applicationService *ApplicationService
	deploymentService  *DeploymentService
}

func NewRepositoryService(dao *daos.Dao, appService *ApplicationService) *RepositoryService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	repositoryService := new(RepositoryService)

	repositoryService.dao = dao
	repositoryService.applicationService = appService

	return repositoryService
}

func (rs *RepositoryService) SyncRepositoryFromDisk(n *tree.Node, record *models.Record) (*models.Record, []*models.Record, error) {

	appRecordIds, err := rs.applicationService.SyncApplicationsFromDisk(n, record)
	if err != nil {
		return nil, nil, apis.NewApiError(500, err.Error(), nil)
	}

	// this complex loop is necessary because we need to know which deployments must add the repository
	// which is currently created, when the creation of the repository is finished.
	// TODO for a future refactoring: extract logic to specific services.
	deployments := make([]*models.Record, 0)
	for _, applicationRecordId := range appRecordIds {

		appRecord, err := rs.dao.FindRecordById(consts.TABLE_APPLICATIONS_NAME, applicationRecordId)
		if err != nil {
			return nil, nil, err
		}

		stagesIds := appRecord.Get(consts.TABLE_APPLICATIONS_FIELD_STAGES).([]string)
		for _, stageId := range stagesIds {

			stageRec, err := rs.dao.FindRecordById(consts.TABLE_STAGES_NAME, stageId)
			if err != nil {
				return nil, nil, err
			}

			deploymentIds := stageRec.Get(consts.TABLE_STAGES_FIELD_DEPLOYMENTS).([]string)
			for _, deploymentId := range deploymentIds {

				deploymentRec, err := rs.dao.FindRecordById(consts.TABLE_DEPLOYMENTS_NAME, deploymentId)
				if err != nil {
					return nil, nil, err
				}

				deployments = append(deployments, deploymentRec)
			}
		}
	}

	return record, deployments, nil
}

func (rs *RepositoryService) FindForName(name string) (*models.Record, error) {

	recs, err := rs.dao.FindRecordsByExpr(consts.TABLE_REPOSITORIES_NAME, dbx.NewExp(consts.TABLE_REPOSITORIES_FIELD_NAME+" = {:"+consts.TABLE_REPOSITORIES_FIELD_NAME+"}", dbx.Params{consts.TABLE_REPOSITORIES_FIELD_NAME: name}))
	if err != nil {
		return nil, err
	}

	if len(recs) > 1 {
		fmt.Println("found more than one entry for repository name. this should not happen.")
	}

	return recs[0], nil
}
