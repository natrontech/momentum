migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "0zk57ok0",
    "name": "parentApplication",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "wf40hpyi2wvpb7y",
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
  collection.schema.removeField("0zk57ok0")

  return dao.saveCollection(collection)
})
