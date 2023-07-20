migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("j5um82igpugrgi2")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "2ns9lvvu",
    "name": "status",
    "type": "select",
    "required": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "UP",
        "ERROR",
        "SYNC"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("j5um82igpugrgi2")

  // remove
  collection.schema.removeField("2ns9lvvu")

  return dao.saveCollection(collection)
})
