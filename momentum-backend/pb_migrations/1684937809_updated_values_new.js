migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  collection.name = "keyValues"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq")

  collection.name = "values_new"

  return dao.saveCollection(collection)
})
