migrate((db) => {
  const collection = new Collection({
    "id": "f8w5oambwthngxo",
    "created": "2023-05-24 13:26:55.408Z",
    "updated": "2023-05-24 13:26:55.408Z",
    "name": "stages",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "lb6cvc8x",
        "name": "name",
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
  const collection = dao.findCollectionByNameOrId("f8w5oambwthngxo");

  return dao.deleteCollection(collection);
})
