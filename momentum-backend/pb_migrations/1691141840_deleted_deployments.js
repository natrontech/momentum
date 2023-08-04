migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "ka4cvffqmuxczw7",
    "created": "2023-05-24 13:40:54.615Z",
    "updated": "2023-05-24 15:24:23.218Z",
    "name": "deployments",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ey4kr9if",
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
        "id": "izaeo3yk",
        "name": "description",
        "type": "text",
        "required": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "seb6k5ba",
        "name": "repositories",
        "type": "relation",
        "required": true,
        "unique": false,
        "options": {
          "collectionId": "os5ld33mgj3dj7b",
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
      },
      {
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
      }
    ],
    "indexes": [
      "CREATE UNIQUE INDEX `idx_1CGxIyJ` ON `deployments` (\n  `keyValues`,\n  `secretKeyValues`\n)",
      "CREATE UNIQUE INDEX `idx_3uvZmVx` ON `deployments` (`keyValues`)",
      "CREATE UNIQUE INDEX `idx_5PfIsEo` ON `deployments` (`secretKeyValues`)"
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
