migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "seb6k5ba",
    "name": "repositories",
    "type": "relation",
    "required": true,
    "unique": false,
    "options": {
      "collectionId": "os5ld33mgj3dj7b",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": [
        "name"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // remove
  collection.schema.removeField("seb6k5ba")

  return dao.saveCollection(collection)
})
