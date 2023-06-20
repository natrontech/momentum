package momentumservices

import (
	consts "momentum/momentum-core/momentum-config"

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

func (kvs *KeyValueService) GetKeyValueCollection() (*models.Collection, error) {

	coll, err := kvs.dao.FindCollectionByNameOrId(consts.TABLE_KEYVALUE_NAME)
	if err != nil {
		return nil, err
	}

	return coll, nil
}

func (kvs *KeyValueService) saveWithoutEvent(record *models.Record) error {
	return kvs.dao.Clone().SaveRecord(record)
}
