migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "r3jb5rhz",
    "name": "status",
    "type": "select",
    "required": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "PENDING",
        "SYNCING",
        "UP-TO-DATE"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "r3jb5rhz",
    "name": "status",
    "type": "select",
    "required": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "PENDING",
        "SYNCING",
        "SYNCED"
      ]
    }
  }))

  return dao.saveCollection(collection)
})
