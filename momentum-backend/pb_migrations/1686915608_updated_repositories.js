migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "nmmegt3m",
    "name": "applications",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "wf40hpyi2wvpb7y",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // remove
  collection.schema.removeField("nmmegt3m")

  return dao.saveCollection(collection)
})
