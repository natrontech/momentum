migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "j9qxbcsd",
    "name": "keyValues",
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
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // remove
  collection.schema.removeField("j9qxbcsd")

  return dao.saveCollection(collection)
})
