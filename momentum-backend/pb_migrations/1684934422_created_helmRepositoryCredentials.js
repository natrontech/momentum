migrate((db) => {
  const collection = new Collection({
    "id": "v81b0cv9ghqpmjr",
    "created": "2023-05-24 13:20:22.475Z",
    "updated": "2023-05-24 13:20:22.475Z",
    "name": "helmRepositoryCredentials",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "smqdrvj0",
        "name": "username",
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
        "id": "o2fcjala",
        "name": "password",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("v81b0cv9ghqpmjr");

  return dao.deleteCollection(collection);
})
