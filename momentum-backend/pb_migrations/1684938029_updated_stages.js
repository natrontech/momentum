migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)",
    "CREATE INDEX `idx_skQyJ7x` ON `stages` (`secretKeyValues`)"
  ]

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "xirmquyd",
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
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)"
  ]

  // remove
  collection.schema.removeField("xirmquyd")

  return dao.saveCollection(collection)
})
