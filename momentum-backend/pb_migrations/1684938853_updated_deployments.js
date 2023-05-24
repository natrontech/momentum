migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "qmvzhwkm",
    "name": "keyValues",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "zp90bz3osxtcevq",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": [
        "key",
        "value"
      ]
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vjhter9p",
    "name": "secretKeyValues",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "ujpm7pjc0i3qg81",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": [
        "key",
        "value"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "qmvzhwkm",
    "name": "keyValues",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "zp90bz3osxtcevq",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": []
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vjhter9p",
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
})
