package momentumservices

import (
	"errors"
	"fmt"
	model "momentum/momentum-core/momentum-model"
	tree "momentum/momentum-core/momentum-tree"
	"strings"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type RepositorySyncService struct {
	dao                *daos.Dao
	applicationService *ApplicationService
	stageService       *StageService
	deploymentService  *DeploymentService
	keyValueService    *KeyValueService
}

func NewRepositorySyncService(
	dao *daos.Dao,
	applicationService *ApplicationService,
	stageService *StageService,
	deploymentService *DeploymentService,
	keyValueService *KeyValueService) *RepositorySyncService {

	repoSyncService := new(RepositorySyncService)
	repoSyncService.dao = dao
	repoSyncService.applicationService = applicationService
	repoSyncService.stageService = stageService
	repoSyncService.deploymentService = deploymentService
	repoSyncService.keyValueService = keyValueService

	return repoSyncService
}

func (rs *RepositorySyncService) SyncRepositoryFromDisk(n *tree.Node, record *models.Record) (*models.Record, []*models.Record, []*models.Record, error) {

	appRecords, err := rs.SyncApplicationsFromDisk(n, record)
	if err != nil {
		return nil, nil, nil, apis.NewApiError(500, err.Error(), nil)
	}

	appRecIds := make([]string, 0)
	for _, appRec := range appRecords {
		appRecIds = append(appRecIds, appRec.Id)
	}
	record.Set(model.TABLE_REPOSITORIES_FIELD_APPLICATIONS, appRecIds)

	// this complex loop is necessary because we need to know which deployments must add the repository
	// which is currently created, when the creation of the repository is finished.
	// TODO for a future refactoring: extract logic to specific services.
	deployments := make([]*models.Record, 0)
	for _, applicationRecord := range appRecords {

		appRecord, err := rs.dao.FindRecordById(model.TABLE_APPLICATIONS_NAME, applicationRecord.Id)
		if err != nil {
			return nil, nil, nil, err
		}

		stagesIds := appRecord.Get(model.TABLE_APPLICATIONS_FIELD_STAGES).([]string)
		for _, stageId := range stagesIds {

			stageRec, err := rs.dao.FindRecordById(model.TABLE_STAGES_NAME, stageId)
			if err != nil {
				return nil, nil, nil, err
			}

			deploymentIds := stageRec.Get(model.TABLE_STAGES_FIELD_DEPLOYMENTS).([]string)
			for _, deploymentId := range deploymentIds {

				deploymentRec, err := rs.dao.FindRecordById(model.TABLE_DEPLOYMENTS_NAME, deploymentId)
				if err != nil {
					return nil, nil, nil, err
				}

				deployments = append(deployments, deploymentRec)
			}
		}
	}

	return record, appRecords, deployments, nil
}

func (rs *RepositorySyncService) SyncApplicationsFromDisk(n *tree.Node, record *models.Record) ([]*models.Record, error) {

	recs := make([]*models.Record, 0)
	apps := n.Apps()
	for _, a := range apps {
		fmt.Println(a)
	}
	for _, app := range apps {

		stages, err := rs.SyncStagesFromDisk(app)
		if err != nil {
			return nil, err
		}

		stageIds := make([]string, 0)
		for _, stage := range stages {
			stageIds = append(stageIds, stage.Id)
		}

		rec, err := rs.applicationService.createWithoutEvent(app.NormalizedPath(), stageIds)
		if err != nil {
			return nil, err
		}

		err = rs.stageService.AddParentApplication(stageIds, rec)

		recs = append(recs, rec)
	}
	return recs, nil
}

func (rs *RepositorySyncService) AddRepository(repositoryRecord *models.Record, applications []*models.Record) error {

	if repositoryRecord.Collection().Name != model.TABLE_REPOSITORIES_NAME {
		return errors.New("repositoryRecord is not record of repositories collection")
	}

	for _, app := range applications {

		app.Set(model.TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY, repositoryRecord.Id)
		err := rs.saveWithoutEvent(app)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rs *RepositorySyncService) SyncStagesFromDisk(n *tree.Node) ([]*models.Record, error) {

	stages := n.AllStages()
	stageRecords := make([]*models.Record, 0)
	var lastStageNode *tree.Node = nil
	var lastStage *models.Record = nil
	for _, stage := range stages {

		deployments, err := rs.SyncDeploymentsFromDisk(n)
		if err != nil {
			return nil, err
		}

		stageRecord, err := rs.stageService.CreateWithoutEvent(stage.NormalizedPath(), deployments)
		if err != nil {
			return nil, err
		}

		if stage.Kind == tree.Directory {
			stageFiles := stage.Files()
			for _, f := range stageFiles {

				err := rs.SyncFile(f, stageRecord)
				if err != nil {
					return nil, err
				}
			}
		}

		err = rs.deploymentService.AddParentStage(stageRecord, deployments)
		if err != nil {
			return nil, err
		}

		if lastStage != nil && lastStageNode != nil && stage.Parent != nil && stage.Parent.IsStage() && lastStageNode.FullPath() == stage.Parent.FullPath() {
			err = rs.stageService.addParentStage(lastStage, stageRecord)
			if err != nil {
				return nil, err
			}
		}

		stageRecords = append(stageRecords, stageRecord)
		lastStage = stageRecord
		lastStageNode = stage
	}

	return stageRecords, nil
}

func (rs *RepositorySyncService) SyncDeploymentsFromDisk(n *tree.Node) ([]*models.Record, error) {

	deployments := n.AllDeployments()

	deploymentIds := make([]*models.Record, 0)
	for _, deployment := range deployments {

		deploymentRecord, err := rs.deploymentService.createWithoutEvent(deployment.NormalizedPath())
		if err != nil {
			return nil, err
		}

		if deployment.Kind == tree.File {

			err := rs.SyncFile(deployment, deploymentRecord)
			if err != nil {
				return nil, err
			}
		}

		deploymentIds = append(deploymentIds, deploymentRecord)
	}

	return deploymentIds, nil
}

func (rs *RepositorySyncService) SyncFile(n *tree.Node, parentArtifact *models.Record) error {

	if n.Kind != tree.File {
		return errors.New("can only sync nodes of type file")
	}

	return rs.syncChildren(n.Children, parentArtifact, n.NormalizedPath())
}

func (rs *RepositorySyncService) syncChildren(children []*tree.Node, parentArtifact *models.Record, filename string) error {

	var err error = nil

	for _, child := range children {

		if len(child.Children) > 0 {

			rs.syncChildren(child.Children, parentArtifact, filename)
		} else {

			if child.Value == "" {
				fmt.Println("empty leaf at:", child.FullPath())
				break
			}

			kvColl, err := rs.keyValueService.GetKeyValueCollection()
			if err != nil {
				break
			}

			propertyPath := strings.Split(child.FullPath(), filename)[1]

			childRecord := models.NewRecord(kvColl)
			childRecord.Set(model.TABLE_KEYVALUE_FIELD_KEY, propertyPath)
			childRecord.Set(model.TABLE_KEYVALUE_FIELD_VALUE, child.Value)
			childRecord.Set(model.TABLE_KEYVALUE_FIELD_DISPLAY_NAME, child.NormalizedPath())

			err = rs.saveWithoutEvent(childRecord)
			if err != nil {
				break
			}

			err = rs.addParentArtifact(parentArtifact, childRecord)
			if err != nil {
				break
			}

			currentKeyValues, ok := parentArtifact.Get(model.TABLE_DEPLOYMENTS_FIELD_KEYVALUES).([]string)
			if ok {
				parentArtifact.Set(model.TABLE_DEPLOYMENTS_FIELD_KEYVALUES, append(currentKeyValues, childRecord.Id))
			} else {
				parentArtifact.Set(model.TABLE_DEPLOYMENTS_FIELD_KEYVALUES, childRecord.Id)
			}
			err = rs.saveWithoutEvent(parentArtifact)
			if err != nil {
				break
			}
		}
	}

	return err
}

func (rs *RepositorySyncService) addParentArtifact(parentArtifact *models.Record, keyValues *models.Record) error {

	switch parentArtifact.Collection().Name {
	case model.TABLE_STAGES_NAME:
		return rs.addParentStage(parentArtifact, []*models.Record{keyValues})
	case model.TABLE_DEPLOYMENTS_NAME:
		return rs.addParentDeployment(parentArtifact, []*models.Record{keyValues})
	default:
		return errors.New("invalid parent record type")
	}
}

func (rs *RepositorySyncService) addParentStage(stage *models.Record, keyValues []*models.Record) error {

	if stage.Collection().Name != model.TABLE_STAGES_NAME {
		return errors.New("parent stage must be record of collection stages")
	}

	for _, kv := range keyValues {

		if kv.Collection().Name != model.TABLE_KEYVALUE_NAME {
			return errors.New("expected keyvalues record type to add parent stage")
		}

		kv.Set(model.TABLE_KEYVALUE_FIELD_PARENTSTAGE, stage.Id)
		err := rs.saveWithoutEvent(kv)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rs *RepositorySyncService) addParentDeployment(deployment *models.Record, keyValues []*models.Record) error {

	if deployment.Collection().Name != model.TABLE_DEPLOYMENTS_NAME {
		return errors.New("parent deployment must be record of collection deploments")
	}

	for _, kv := range keyValues {

		if kv.Collection().Name != model.TABLE_KEYVALUE_NAME {
			return errors.New("expected keyvalues record type to add parent deployment")
		}

		kv.Set(model.TABLE_KEYVALUE_FIELD_PARENTDEPLOYMENT, deployment.Id)
		err := rs.saveWithoutEvent(kv)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rs *RepositorySyncService) saveWithoutEvent(record *models.Record) error {
	return rs.dao.Clone().SaveRecord(record)
}
