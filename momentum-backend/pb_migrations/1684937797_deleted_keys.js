migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("5jzvmtxgntdyge5");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "5jzvmtxgntdyge5",
    "created": "2023-05-24 13:40:52.934Z",
    "updated": "2023-05-24 13:53:11.744Z",
    "name": "keys",
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
        "id": "mvji3fg0",
        "name": "parentKey",
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
      },
      {
        "system": false,
        "id": "uhnotvkb",
        "name": "values",
        "type": "relation",
        "required": false,
        "unique": false,
        "options": {
          "collectionId": "wz8leflqf6p8lct",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": null,
          "displayFields": [
            "value"
          ]
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
})
