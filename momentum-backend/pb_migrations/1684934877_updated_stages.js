migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "nyx3xebq",
    "name": "parentStage",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "f8w5oambwthngxo",
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
  collection.schema.removeField("nyx3xebq")

  return dao.saveCollection(collection)
})
