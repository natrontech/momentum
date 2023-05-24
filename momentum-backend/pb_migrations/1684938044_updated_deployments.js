migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_1CGxIyJ` ON `deployments` (\n  `keyValues`,\n  `secretKeyValues`\n)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_1CGxIyJ` ON `deployments` (`keyValues`)"
  ]

  return dao.saveCollection(collection)
})
