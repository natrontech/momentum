migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)",
    "CREATE UNIQUE INDEX `idx_skQyJ7x` ON `stages` (`secretKeyValues`)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)",
    "CREATE INDEX `idx_skQyJ7x` ON `stages` (`secretKeyValues`)"
  ]

  return dao.saveCollection(collection)
})
