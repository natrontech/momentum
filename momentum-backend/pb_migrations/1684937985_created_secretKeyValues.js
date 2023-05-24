migrate((db) => {
  const collection = new Collection({
    "id": "ujpm7pjc0i3qg81",
    "created": "2023-05-24 14:19:45.983Z",
    "updated": "2023-05-24 14:19:45.983Z",
    "name": "secretKeyValues",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "3pvublr5",
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
        "id": "ol84kw30",
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
  const collection = dao.findCollectionByNameOrId("ujpm7pjc0i3qg81");

  return dao.deleteCollection(collection);
})
