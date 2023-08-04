migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "f8w5oambwthngxo",
    "created": "2023-05-24 13:26:55.408Z",
    "updated": "2023-05-24 14:33:55.658Z",
    "name": "stages",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "lb6cvc8x",
        "name": "name",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
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
      },
      {
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
      },
      {
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
      },
      {
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
      },
      {
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
      }
    ],
    "indexes": [
      "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)",
      "CREATE UNIQUE INDEX `idx_skQyJ7x` ON `stages` (`secretKeyValues`)",
      "CREATE UNIQUE INDEX `idx_JXT6tA7` ON `stages` (\n  `secretKeyValues`,\n  `keyValues`\n)"
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
})
