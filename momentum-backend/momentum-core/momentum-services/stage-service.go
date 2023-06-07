package momentumservices

import (
	"fmt"
	tree "momentum/momentum-core/momentum-tree"

	"github.com/pocketbase/pocketbase/daos"
)

type StageService struct {
	dao               *daos.Dao
	deplyomentService *DeplyomentService
}

func NewStageService(dao *daos.Dao, deploymentService *DeplyomentService) *StageService {

	stageService := new(StageService)
	stageService.deplyomentService = deploymentService
	stageService.dao = dao

	return stageService
}

func (ss *StageService) AddStages(n *tree.Node) error {

	stages, err := n.Stages()
	if err != nil {
		return err
	}

	for _, stage := range stages {

		fmt.Println(stage.Path)

		err = ss.AddStages(stage) // Stages are multilevel
		if err != nil {
			return err
		}

		err = ss.deplyomentService.AddDeployments(n)
		if err != nil {
			return err
		}
	}

	return nil
}
