package momentumservices

import (
	"fmt"
	tree "momentum/momentum-core/momentum-tree"

	"github.com/pocketbase/pocketbase/daos"
)

type DeplyomentService struct {
	dao *daos.Dao
}

func NewDeploymentService(dao *daos.Dao) *DeplyomentService {

	deplyomentService := new(DeplyomentService)
	deplyomentService.dao = dao

	return deplyomentService
}

func (ds *DeplyomentService) AddDeployments(n *tree.Node) error {

	deployments, err := n.Deployments()
	if err != nil {
		return err
	}

	for _, deployment := range deployments {

		fmt.Println("Deployment:", deployment.Path)
	}

	return nil
}
