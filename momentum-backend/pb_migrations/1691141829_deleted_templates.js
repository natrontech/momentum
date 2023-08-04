migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("9gf5p3rlwfhgh1j");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "9gf5p3rlwfhgh1j",
    "created": "2023-05-24 14:28:14.243Z",
    "updated": "2023-05-24 14:32:50.565Z",
    "name": "templates",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "5c1k4g8q",
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
        "id": "r5t5pphi",
        "name": "description",
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
      },
      {
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
      }
    ],
    "indexes": [
      "CREATE UNIQUE INDEX `idx_sAYJ45f` ON `templates` (`keyValues`)",
      "CREATE UNIQUE INDEX `idx_X6weBCl` ON `templates` (`secretKeyValues`)",
      "CREATE UNIQUE INDEX `idx_hNLbnBO` ON `templates` (\n  `keyValues`,\n  `secretKeyValues`\n)"
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
