migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5")

  // remove
  collection.schema.removeField("sfq4uvov")

  // remove
  collection.schema.removeField("osbrs5bs")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "sfq4uvov",
    "name": "value",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "osbrs5bs",
    "name": "values",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "5jzvmtxgntdyge5",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
})
