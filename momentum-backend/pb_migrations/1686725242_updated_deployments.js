migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "eb0jvo3o",
    "name": "status",
    "type": "select",
    "required": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "RUNNING",
        "FAILED",
        "PENDING",
        "UNKNOWN"
      ]
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "x8gggfpu",
    "name": "namespace",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // remove
  collection.schema.removeField("eb0jvo3o")

  // remove
  collection.schema.removeField("x8gggfpu")

  return dao.saveCollection(collection)
})
