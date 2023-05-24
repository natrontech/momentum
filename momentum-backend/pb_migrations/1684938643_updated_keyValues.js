migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "yzpsh3rf",
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
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  // remove
  collection.schema.removeField("yzpsh3rf")

  return dao.saveCollection(collection)
})
