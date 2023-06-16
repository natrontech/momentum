package momentumservices

import (
	"errors"
	"fmt"
	consts "momentum/momentum-core/momentum-config"
	tree "momentum/momentum-core/momentum-tree"
	"strings"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type KeyValueService struct {
	dao *daos.Dao
}

func NewKeyValueService(dao *daos.Dao) *KeyValueService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	keyValueService := new(KeyValueService)

	keyValueService.dao = dao

	return keyValueService
}

func (kvs *KeyValueService) SyncFile(n *tree.Node, parentArtifact *models.Record) error {

	if n.Kind != tree.File {
		return errors.New("can only sync nodes of type file")
	}

	return kvs.syncChildren(n.Children, parentArtifact, n.NormalizedPath())
}

func (kvs *KeyValueService) GetKeyValueCollection() (*models.Collection, error) {

	coll, err := kvs.dao.FindCollectionByNameOrId(consts.TABLE_KEYVALUE_NAME)
	if err != nil {
		return nil, err
	}

	return coll, nil
}

func (kvs *KeyValueService) syncChildren(children []*tree.Node, parentArtifact *models.Record, filename string) error {

	var err error = nil

	for _, child := range children {

		if len(child.Children) > 0 {

			kvs.syncChildren(child.Children, parentArtifact, filename)
		} else {

			if child.Value == "" {
				fmt.Println("empty leaf at:", child.FullPath())
				break
			}

			kvColl, err := kvs.GetKeyValueCollection()
			if err != nil {
				break
			}

			propertyPath := strings.Split(child.FullPath(), filename)[1]

			childRecord := models.NewRecord(kvColl)
			childRecord.Set(consts.TABLE_KEYVALUE_FIELD_KEY, propertyPath)
			childRecord.Set(consts.TABLE_KEYVALUE_FIELD_VALUE, child.Value)

			err = kvs.saveWithoutEvent(childRecord)
			if err != nil {
				break
			}

			err = kvs.addParentArtifact(parentArtifact, childRecord)
			if err != nil {
				break
			}

			currentKeyValues, ok := parentArtifact.Get(consts.GENERIC_FIELD_KEYVALUES).([]string)
			if ok {
				parentArtifact.Set(consts.GENERIC_FIELD_KEYVALUES, append(currentKeyValues, childRecord.Id))
			} else {
				parentArtifact.Set(consts.GENERIC_FIELD_KEYVALUES, childRecord.Id)
			}
			err = kvs.saveWithoutEvent(parentArtifact)
			if err != nil {
				break
			}
		}
	}

	return err
}

func (kvs *KeyValueService) addParentArtifact(parentArtifact *models.Record, keyValues *models.Record) error {

	switch parentArtifact.Collection().Name {
	case consts.TABLE_STAGES_NAME:
		return kvs.addParentStage(parentArtifact, []*models.Record{keyValues})
	case consts.TABLE_DEPLOYMENTS_NAME:
		return kvs.addParentDeployment(parentArtifact, []*models.Record{keyValues})
	default:
		return errors.New("invalid parent record type")
	}
}

func (kvs *KeyValueService) addParentStage(stage *models.Record, keyValues []*models.Record) error {

	if stage.Collection().Name != consts.TABLE_STAGES_NAME {
		return errors.New("parent stage must be record of collection stages")
	}

	for _, kv := range keyValues {

		if kv.Collection().Name != consts.TABLE_KEYVALUE_NAME {
			return errors.New("expected keyvalues record type to add parent stage")
		}

		kv.Set(consts.TABLE_KEYVALUE_FIELD_PARENTSTAGE, stage.Id)
		err := kvs.saveWithoutEvent(kv)
		if err != nil {
			return err
		}
	}

	return nil
}

func (kvs *KeyValueService) addParentDeployment(deployment *models.Record, keyValues []*models.Record) error {

	if deployment.Collection().Name != consts.TABLE_DEPLOYMENTS_NAME {
		return errors.New("parent deployment must be record of collection deploments")
	}

	for _, kv := range keyValues {

		if kv.Collection().Name != consts.TABLE_KEYVALUE_NAME {
			return errors.New("expected keyvalues record type to add parent deployment")
		}

		kv.Set(consts.TABLE_KEYVALUE_FIELD_PARENTDEPLOYMENT, deployment.Id)
		err := kvs.saveWithoutEvent(kv)
		if err != nil {
			return err
		}
	}

	return nil
}

func (kvs *KeyValueService) saveWithoutEvent(record *models.Record) error {
	return kvs.dao.Clone().SaveRecord(record)
}
