migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "uhnotvkb",
    "name": "values",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "wz8leflqf6p8lct",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5")

  // remove
  collection.schema.removeField("uhnotvkb")

  return dao.saveCollection(collection)
})
