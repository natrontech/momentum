migrate((db) => {
  const collection = new Collection({
    "id": "wz8leflqf6p8lct",
    "created": "2023-05-24 13:52:07.057Z",
    "updated": "2023-05-24 13:52:07.057Z",
    "name": "values",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "eyf7tx7n",
        "name": "value",
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
  const collection = dao.findCollectionByNameOrId("wz8leflqf6p8lct");

  return dao.deleteCollection(collection);
})
