migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // add
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
        "SYNC",
        "SYNCING",
        "SYNCED"
      ]
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "b43e5pqo",
    "name": "repositoryCredentials",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "epya5jownu486y2",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // remove
  collection.schema.removeField("r3jb5rhz")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "b43e5pqo",
    "name": "field",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "epya5jownu486y2",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
})
