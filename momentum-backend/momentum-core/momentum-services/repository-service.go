package momentumservices

import (
	"fmt"
	model "momentum/momentum-core/momentum-model"

	"github.com/pocketbase/dbx"
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

func (rs *RepositoryService) FindForName(name string) (*models.Record, error) {

	recs, err := rs.dao.FindRecordsByExpr(model.TABLE_REPOSITORIES_NAME, dbx.NewExp(model.TABLE_REPOSITORIES_FIELD_NAME+" = {:"+model.TABLE_REPOSITORIES_FIELD_NAME+"}", dbx.Params{model.TABLE_REPOSITORIES_FIELD_NAME: name}))
	if err != nil {
		return nil, err
	}

	if len(recs) > 1 {
		fmt.Println("found more than one entry for repository name. this should not happen.")
	}

	return recs[0], nil
}
