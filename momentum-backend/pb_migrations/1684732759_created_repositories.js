migrate((db) => {
  const collection = new Collection({
    "id": "os5ld33mgj3dj7b",
    "created": "2023-05-22 05:19:19.127Z",
    "updated": "2023-05-22 05:19:19.127Z",
    "name": "repositories",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ipzx56fk",
        "name": "name",
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
        "id": "cxcgjxz4",
        "name": "url",
        "type": "url",
        "required": false,
        "unique": false,
        "options": {
          "exceptDomains": [],
          "onlyDomains": []
        }
      },
      {
        "system": false,
        "id": "b43e5pqo",
        "name": "field",
        "type": "relation",
        "required": false,
        "unique": false,
        "options": {
          "collectionId": "epya5jownu486y2",
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
  const collection = dao.findCollectionByNameOrId("os5ld33mgj3dj7b");

  return dao.deleteCollection(collection);
})
