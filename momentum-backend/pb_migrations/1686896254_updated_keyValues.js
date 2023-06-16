migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "tmctz4dv",
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

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "alxyew1k",
    "name": "parentDeployment",
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
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  // remove
  collection.schema.removeField("tmctz4dv")

  // remove
  collection.schema.removeField("alxyew1k")

  return dao.saveCollection(collection)
})
