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

	kvs.syncChildren(n.Children, parentArtifact, n.NormalizedPath())

	return nil
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

			err = kvs.dao.Clone().SaveRecord(childRecord)
			if err != nil {
				break
			}

			currentKeyValues, ok := parentArtifact.Get(consts.GENERIC_FIELD_KEYVALUES).([]string)
			if ok {
				parentArtifact.Set(consts.GENERIC_FIELD_KEYVALUES, append(currentKeyValues, childRecord.Id))
			} else {
				parentArtifact.Set(consts.GENERIC_FIELD_KEYVALUES, childRecord.Id)
			}
			err = kvs.dao.Clone().SaveRecord(parentArtifact)
			if err != nil {
				break
			}
		}
	}

	return err
}
