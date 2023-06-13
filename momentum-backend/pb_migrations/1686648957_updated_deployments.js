migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  collection.indexes = []

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_1CGxIyJ` ON `deployments` (\n  `keyValues`,\n  `secretKeyValues`\n)",
    "CREATE UNIQUE INDEX `idx_3uvZmVx` ON `deployments` (`keyValues`)",
    "CREATE UNIQUE INDEX `idx_5PfIsEo` ON `deployments` (`secretKeyValues`)"
  ]

  return dao.saveCollection(collection)
})
