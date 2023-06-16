migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("wf40hpyi2wvpb7y")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "cc62mdka",
    "name": "parentRepository",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "os5ld33mgj3dj7b",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("wf40hpyi2wvpb7y")

  // remove
  collection.schema.removeField("cc62mdka")

  return dao.saveCollection(collection)
})
