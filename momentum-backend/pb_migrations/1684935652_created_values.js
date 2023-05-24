migrate((db) => {
  const collection = new Collection({
    "id": "5jzvmtxgntdyge5",
    "created": "2023-05-24 13:40:52.934Z",
    "updated": "2023-05-24 13:40:52.934Z",
    "name": "values",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "l3dhxtgz",
        "name": "key",
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
        "id": "sfq4uvov",
        "name": "value",
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
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5");

  return dao.deleteCollection(collection);
})
