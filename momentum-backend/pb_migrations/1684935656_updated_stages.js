migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // add
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
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // remove
  collection.schema.removeField("z5xx0qqt")

  return dao.saveCollection(collection)
})
