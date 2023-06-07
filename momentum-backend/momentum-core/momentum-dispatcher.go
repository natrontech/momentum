package momentumcore

import (
	"fmt"

	momentumconfig "momentum/momentum-core/momentum-config"
	momentumcontrollers "momentum/momentum-core/momentum-controllers"
	momentumservices "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

type MomentumDispatcherRuleType int

const (
	CREATE MomentumDispatcherRuleType = 1 << iota
	UPDATE
	DELETE
)

type MomentumDispatcherRule struct {
	tableName string
	action    func(*models.Record, *momentumconfig.MomentumConfig) error
}

type MomentumDispatcher struct {
	App    *pocketbase.PocketBase
	Config *momentumconfig.MomentumConfig

	CreateRules []*MomentumDispatcherRule
	UpdateRules []*MomentumDispatcherRule
	DeleteRules []*MomentumDispatcherRule

	RepositoryController   *momentumcontrollers.RepositoryController
	ApplicationsController *momentumcontrollers.ApplicationController
	StagesController       *momentumcontrollers.StageController
	DeploymentController   *momentumcontrollers.DeploymentController
}

func NewDispatcher(config *momentumconfig.MomentumConfig, app *pocketbase.PocketBase) *MomentumDispatcher {

	dispatcher := new(MomentumDispatcher)
	dispatcher.Config = config

	deploymentService := momentumservices.NewDeploymentService(app.Dao())
	stageService := momentumservices.NewStageService(app.Dao(), deploymentService)
	appService := momentumservices.NewApplicationService(app.Dao(), stageService)

	dispatcher.RepositoryController = momentumcontrollers.NewRepositoryController(appService)
	dispatcher.ApplicationsController = momentumcontrollers.NewApplicationController(appService)
	dispatcher.StagesController = momentumcontrollers.NewStageController(stageService)
	dispatcher.DeploymentController = momentumcontrollers.NewDeploymentController(deploymentService)

	dispatcher.CreateRules = dispatcher.setupCreateRules()
	dispatcher.UpdateRules = dispatcher.setupUpdateRules()
	dispatcher.DeleteRules = dispatcher.setupDeleteRules()

	dispatcher.App = app

	return dispatcher
}

func (d *MomentumDispatcher) DispatchCreate(recordEvent *core.RecordCreateEvent) error {

	for _, rule := range d.CreateRules {
		fmt.Println("Rule:", rule.tableName)
		if rule.tableName == recordEvent.Record.TableName() {
			rule.action(recordEvent.Record, d.Config)
		}
	}
	return nil
}

func (d *MomentumDispatcher) DispatchUpdate(recordEvent *core.RecordUpdateEvent) error {

	for _, rule := range d.UpdateRules {
		if rule.tableName == recordEvent.Record.TableName() {
			rule.action(recordEvent.Record, d.Config)
		}
	}
	return nil
}

func (d *MomentumDispatcher) DispatchDelete(recordEvent *core.RecordDeleteEvent) error {

	for _, rule := range d.DeleteRules {
		if rule.tableName == recordEvent.Record.TableName() {
			rule.action(recordEvent.Record, d.Config)
		}
	}
	return nil
}

func (d *MomentumDispatcher) setupCreateRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{momentumconfig.TABLE_REPOSITORIES_NAME, d.RepositoryController.AddRepository},
		{momentumconfig.TABLE_APPLICATIONS_NAME, d.ApplicationsController.AddApplication},
		{momentumconfig.TABLE_STAGES_NAME, d.StagesController.AddStage},
		{momentumconfig.TABLE_DEPLOYMENTS_NAME, d.DeploymentController.AddDeplyoment},
	}
}

func (d *MomentumDispatcher) setupUpdateRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{momentumconfig.TABLE_REPOSITORIES_NAME, d.RepositoryController.UpdateRepository},
		{momentumconfig.TABLE_APPLICATIONS_NAME, d.ApplicationsController.UpdateApplication},
		{momentumconfig.TABLE_STAGES_NAME, d.StagesController.UpdateStage},
		{momentumconfig.TABLE_DEPLOYMENTS_NAME, d.DeploymentController.UpdateDeplyoment},
	}
}

func (d *MomentumDispatcher) setupDeleteRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{momentumconfig.TABLE_REPOSITORIES_NAME, d.RepositoryController.DeleteRepository},
		{momentumconfig.TABLE_APPLICATIONS_NAME, d.ApplicationsController.DeleteApplication},
		{momentumconfig.TABLE_STAGES_NAME, d.StagesController.DeleteStage},
		{momentumconfig.TABLE_DEPLOYMENTS_NAME, d.DeploymentController.DeleteDeplyoment},
	}
}
