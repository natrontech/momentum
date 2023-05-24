migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "nyx3xebq",
    "name": "parentStage",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "f8w5oambwthngxo",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": [
        "name"
      ]
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "z5xx0qqt",
    "name": "deployments",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "ka4cvffqmuxczw7",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": [
        "name"
      ]
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "djv5hxb6",
    "name": "template",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "9gf5p3rlwfhgh1j",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": [
        "name"
      ]
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "6f4m9lvy",
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
    "id": "xirmquyd",
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
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "nyx3xebq",
    "name": "parentStage",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "f8w5oambwthngxo",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "z5xx0qqt",
    "name": "deployments",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "ka4cvffqmuxczw7",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": []
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "djv5hxb6",
    "name": "template",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "9gf5p3rlwfhgh1j",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "6f4m9lvy",
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
    "id": "xirmquyd",
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
