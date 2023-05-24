migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5")

  // remove
  collection.schema.removeField("pt35odad")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "mvji3fg0",
    "name": "parentKey",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "5jzvmtxgntdyge5",
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

  // remove
  collection.schema.removeField("mvji3fg0")

  return dao.saveCollection(collection)
})
