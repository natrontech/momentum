migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)",
    "CREATE UNIQUE INDEX `idx_skQyJ7x` ON `stages` (`secretKeyValues`)",
    "CREATE UNIQUE INDEX `idx_JXT6tA7` ON `stages` (\n  `secretKeyValues`,\n  `keyValues`\n)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)",
    "CREATE UNIQUE INDEX `idx_skQyJ7x` ON `stages` (`secretKeyValues`)"
  ]

  return dao.saveCollection(collection)
})
