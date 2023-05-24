migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9gf5p3rlwfhgh1j")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "arhbhwoc",
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
        "displayName"
      ]
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "394csawp",
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
        "displayName"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9gf5p3rlwfhgh1j")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "arhbhwoc",
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
    "id": "394csawp",
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
        "key"
      ]
    }
  }))

  return dao.saveCollection(collection)
})
