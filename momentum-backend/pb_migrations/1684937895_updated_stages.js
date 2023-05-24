migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)"
  ]

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "6f4m9lvy",
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

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "z5xx0qqt",
    "name": "deployments",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "ka4cvffqmuxczw7",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = []

  // remove
  collection.schema.removeField("6f4m9lvy")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "z5xx0qqt",
    "name": "deployments",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "ka4cvffqmuxczw7",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
})
