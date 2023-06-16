migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "juvm3tmo",
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
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // remove
  collection.schema.removeField("juvm3tmo")

  return dao.saveCollection(collection)
})
