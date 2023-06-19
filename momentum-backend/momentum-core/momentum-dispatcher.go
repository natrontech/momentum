package momentumcore

import (
	"errors"
	"fmt"

	kustomizeclient "momentum/kustomize-client"
	conf "momentum/momentum-core/momentum-config"
	controllers "momentum/momentum-core/momentum-controllers"
	services "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

var REPOSITORY_ADDED_EVENT_CHANNEL = make(chan *controllers.RepositoryAddedEvent, 1)

type MomentumDispatcherRule struct {
	tableName string
	action    func(*models.Record, *conf.MomentumConfig) error
}

type MomentumDispatcher struct {
	Pocketbase *pocketbase.PocketBase
	Config     *conf.MomentumConfig

	CreateRules []*MomentumDispatcherRule
	UpdateRules []*MomentumDispatcherRule
	DeleteRules []*MomentumDispatcherRule

	RepositoryController   *controllers.RepositoryController
	ApplicationsController *controllers.ApplicationController
	StagesController       *controllers.StageController
	DeploymentController   *controllers.DeploymentController

	kustomizeValidator *kustomizeclient.KustomizationValidationService
}

func NewDispatcher(config *conf.MomentumConfig, pb *pocketbase.PocketBase) *MomentumDispatcher {

	// the order of statements is relevant
	dispatcher := new(MomentumDispatcher)
	dispatcher.Config = config

	keyValueService := services.NewKeyValueService(pb.Dao())
	deploymentService := services.NewDeploymentService(pb.Dao(), keyValueService)
	stageService := services.NewStageService(pb.Dao(), deploymentService, keyValueService)
	appService := services.NewApplicationService(pb.Dao(), stageService)
	repoService := services.NewRepositoryService(pb.Dao(), appService)

	dispatcher.kustomizeValidator = kustomizeclient.NewKustomizationValidationService(dispatcher.Config, repoService)

	dispatcher.RepositoryController = controllers.NewRepositoryController(repoService, deploymentService, REPOSITORY_ADDED_EVENT_CHANNEL, dispatcher.kustomizeValidator)
	dispatcher.ApplicationsController = controllers.NewApplicationController(appService, repoService)
	dispatcher.StagesController = controllers.NewStageController(stageService)
	dispatcher.DeploymentController = controllers.NewDeploymentController(deploymentService, repoService)

	dispatcher.CreateRules = dispatcher.setupCreateRules()
	dispatcher.UpdateRules = dispatcher.setupUpdateRules()
	dispatcher.DeleteRules = dispatcher.setupDeleteRules()

	dispatcher.Pocketbase = pb

	dispatcher.setupRepositoryAddedEventChannelObserver()

	return dispatcher
}

func (d *MomentumDispatcher) DispatchCreate(recordEvent *core.RecordCreateEvent) error {

	for _, rule := range d.CreateRules {
		fmt.Println("Rule:", rule.tableName)
		if rule.tableName == recordEvent.Record.TableName() {
			err := rule.action(recordEvent.Record, d.Config)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *MomentumDispatcher) DispatchUpdate(recordEvent *core.RecordUpdateEvent) error {

	for _, rule := range d.UpdateRules {
		if rule.tableName == recordEvent.Record.TableName() {
			err := rule.action(recordEvent.Record, d.Config)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *MomentumDispatcher) DispatchDelete(recordEvent *core.RecordDeleteEvent) error {

	for _, rule := range d.DeleteRules {
		if rule.tableName == recordEvent.Record.TableName() {
			err := rule.action(recordEvent.Record, d.Config)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *MomentumDispatcher) setupCreateRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{conf.TABLE_REPOSITORIES_NAME, d.RepositoryController.AddRepository},
		{conf.TABLE_APPLICATIONS_NAME, d.ApplicationsController.AddApplication},
		{conf.TABLE_STAGES_NAME, d.StagesController.AddStage},
		{conf.TABLE_DEPLOYMENTS_NAME, d.DeploymentController.AddDeployment},
	}
}

func (d *MomentumDispatcher) setupUpdateRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{conf.TABLE_REPOSITORIES_NAME, d.RepositoryController.UpdateRepository},
		{conf.TABLE_APPLICATIONS_NAME, d.ApplicationsController.UpdateApplication},
		{conf.TABLE_STAGES_NAME, d.StagesController.UpdateStage},
		{conf.TABLE_DEPLOYMENTS_NAME, d.DeploymentController.UpdateDeployment},
	}
}

func (d *MomentumDispatcher) setupDeleteRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{conf.TABLE_REPOSITORIES_NAME, d.RepositoryController.DeleteRepository},
		{conf.TABLE_APPLICATIONS_NAME, d.ApplicationsController.DeleteApplication},
		{conf.TABLE_STAGES_NAME, d.StagesController.DeleteStage},
		{conf.TABLE_DEPLOYMENTS_NAME, d.DeploymentController.DeleteDeployment},
	}
}

func (d *MomentumDispatcher) setupRepositoryAddedEventChannelObserver() {

	d.Pocketbase.OnRecordAfterCreateRequest(conf.TABLE_REPOSITORIES_NAME).Add(func(e *core.RecordCreateEvent) error {

		event := <-REPOSITORY_ADDED_EVENT_CHANNEL

		err := d.DeploymentController.AddRepositoryToDeployments(event)
		if err != nil {
			fmt.Println("failed adding relationship to deployments for repository after receiving RepositoryAddedEvent:", event, err, err.Error())
			return err
		}

		err = d.ApplicationsController.AddRepositoryToApplications(event)
		if err != nil {
			fmt.Println("failed adding relationship to applications for repository after receiving RepositoryAddedEvent:", event, err, err.Error())
			return err
		}

		validationSuccessful, err := d.kustomizeValidator.Validate(event.RepositoryName)
		if !validationSuccessful {
			if err != nil {
				return nil
			}
			return errors.New("validation failed for repository on repository added event")
		}

		return nil
	})
}
