migrate((db) => {
  const collection = new Collection({
    "id": "zp90bz3osxtcevq",
    "created": "2023-05-24 14:12:20.499Z",
    "updated": "2023-05-24 14:12:20.499Z",
    "name": "values_new",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "e23wk1cr",
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
        "id": "z9aaouuo",
        "name": "value",
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
        "id": "b0qfhhmm",
        "name": "type",
        "type": "select",
        "required": false,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "values": [
            "value",
            "array",
            "map"
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
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("zp90bz3osxtcevq");

  return dao.deleteCollection(collection);
})
