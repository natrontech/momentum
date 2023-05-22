migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  collection.listRule = ""
  collection.viewRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b")

  collection.listRule = null
  collection.viewRule = null

  return dao.saveCollection(collection)
})
