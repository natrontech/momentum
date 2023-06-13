migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Awdhj8M` ON `repositories` (`name`)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  collection.indexes = []

  return dao.saveCollection(collection)
})
