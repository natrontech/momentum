migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // remove
  collection.schema.removeField("r3jb5rhz")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "r3jb5rhz",
    "name": "status",
    "type": "select",
    "required": true,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "PENDING",
        "SYNCING",
        "UP-TO-DATE",
        "ERROR"
      ]
    }
  }))

  return dao.saveCollection(collection)
})
