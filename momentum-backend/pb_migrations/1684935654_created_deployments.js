migrate((db) => {
  const collection = new Collection({
    "id": "ka4cvffqmuxczw7",
    "created": "2023-05-24 13:40:54.615Z",
    "updated": "2023-05-24 13:40:54.615Z",
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
        "id": "qfceudkx",
        "name": "values",
        "type": "relation",
        "required": false,
        "unique": false,
        "options": {
          "collectionId": "5jzvmtxgntdyge5",
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
  const collection = dao.findCollectionByNameOrId("ka4cvffqmuxczw7");

  return dao.deleteCollection(collection);
})
