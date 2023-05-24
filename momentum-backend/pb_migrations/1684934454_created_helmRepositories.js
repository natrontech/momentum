migrate((db) => {
  const collection = new Collection({
    "id": "z6e64e4y3ifjd4v",
    "created": "2023-05-24 13:20:54.976Z",
    "updated": "2023-05-24 13:20:54.976Z",
    "name": "helmRepositories",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "lltfaiwh",
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
        "id": "l1vjlw2g",
        "name": "url",
        "type": "url",
        "required": true,
        "unique": false,
        "options": {
          "exceptDomains": [],
          "onlyDomains": []
        }
      },
      {
        "system": false,
        "id": "lfh7brfk",
        "name": "helmRepositoryCredentials",
        "type": "relation",
        "required": false,
        "unique": false,
        "options": {
          "collectionId": "v81b0cv9ghqpmjr",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": []
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
  const collection = dao.findCollectionByNameOrId("z6e64e4y3ifjd4v");

  return dao.deleteCollection(collection);
})
