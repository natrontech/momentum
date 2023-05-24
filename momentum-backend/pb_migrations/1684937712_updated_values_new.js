migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  // remove
  collection.schema.removeField("manapysh")

  // remove
  collection.schema.removeField("hjsxvuxp")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "manapysh",
    "name": "isCollection",
    "type": "bool",
    "required": false,
    "unique": false,
    "options": {}
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "hjsxvuxp",
    "name": "order",
    "type": "number",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null
    }
  }))

  return dao.saveCollection(collection)
})
