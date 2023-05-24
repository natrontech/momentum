migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5")

  collection.name = "keys"

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "pt35odad",
    "name": "parentKeys",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "_pb_users_auth_",
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

  collection.name = "values"

  // remove
  collection.schema.removeField("pt35odad")

  return dao.saveCollection(collection)
})
