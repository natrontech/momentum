package momentumservices

import (
	"fmt"
	tree "momentum/momentum-core/momentum-tree"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type ApplicationService struct {
	dao          *daos.Dao
	stageService *StageService
}

func NewApplicationService(dao *daos.Dao, stageService *StageService) *ApplicationService {

	appService := new(ApplicationService)
	appService.dao = dao
	appService.stageService = stageService

	return appService
}

func (as *ApplicationService) AddApplications(n *tree.Node, record *models.Record) error {

	apps := n.Apps()
	for _, app := range apps {

		fmt.Println(app.Path)

		err := as.stageService.AddStages(app)
		if err != nil {
			return err
		}
	}

	return nil
}
