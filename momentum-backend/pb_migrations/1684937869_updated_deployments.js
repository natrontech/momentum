migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_1CGxIyJ` ON `deployments` (`keyValues`)"
  ]

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "qmvzhwkm",
    "name": "keyValues",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "zp90bz3osxtcevq",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  collection.indexes = []

  // remove
  collection.schema.removeField("qmvzhwkm")

  return dao.saveCollection(collection)
})
