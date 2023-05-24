migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("wz8leflqf6p8lct")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_UryBHoz` ON `values` (`value`)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("wz8leflqf6p8lct")

  collection.indexes = []

  return dao.saveCollection(collection)
})
