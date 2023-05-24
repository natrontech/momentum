migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "djv5hxb6",
    "name": "template",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "9gf5p3rlwfhgh1j",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // remove
  collection.schema.removeField("djv5hxb6")

  return dao.saveCollection(collection)
})
