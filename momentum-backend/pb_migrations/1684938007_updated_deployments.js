migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vjhter9p",
    "name": "secretKeyValues",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "ujpm7pjc0i3qg81",
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
  collection.schema.removeField("vjhter9p")

  return dao.saveCollection(collection)
})
