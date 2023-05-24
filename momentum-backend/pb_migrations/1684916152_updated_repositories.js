migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ipzx56fk",
    "name": "name",
    "type": "text",
    "required": true,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "cxcgjxz4",
    "name": "url",
    "type": "url",
    "required": true,
    "unique": false,
    "options": {
      "exceptDomains": [],
      "onlyDomains": []
    }
  }))

  // update
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
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ipzx56fk",
    "name": "name",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "cxcgjxz4",
    "name": "url",
    "type": "url",
    "required": false,
    "unique": false,
    "options": {
      "exceptDomains": [],
      "onlyDomains": []
    }
  }))

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
        "UP-TO-DATE",
        "ERROR"
      ]
    }
  }))

  return dao.saveCollection(collection)
})
