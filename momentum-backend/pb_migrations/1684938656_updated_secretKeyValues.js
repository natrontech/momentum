migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ujpm7pjc0i3qg81")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "od3uiju3",
    "name": "displayName",
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
  const collection = dao.findCollectionByNameOrId("ujpm7pjc0i3qg81")

  // remove
  collection.schema.removeField("od3uiju3")

  return dao.saveCollection(collection)
})
