migrate((db) => {
  const collection = new Collection({
    "id": "epya5jownu486y2",
    "created": "2023-05-22 05:18:52.469Z",
    "updated": "2023-05-22 05:18:52.469Z",
    "name": "repositoryCredentials",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "gmp9r8re",
        "name": "username",
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
        "id": "y663szrk",
        "name": "password",
        "type": "text",
        "required": false,
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
  const collection = dao.findCollectionByNameOrId("epya5jownu486y2");

  return dao.deleteCollection(collection);
})
