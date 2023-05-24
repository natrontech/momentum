migrate((db) => {
  const collection = new Collection({
    "id": "wf40hpyi2wvpb7y",
    "created": "2023-05-24 13:26:59.748Z",
    "updated": "2023-05-24 13:26:59.748Z",
    "name": "applications",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ytbnpk0r",
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
        "id": "bw7pupoi",
        "name": "helmRepository",
        "type": "relation",
        "required": false,
        "unique": false,
        "options": {
          "collectionId": "z6e64e4y3ifjd4v",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": []
        }
      },
      {
        "system": false,
        "id": "ydrlb6pn",
        "name": "stages",
        "type": "relation",
        "required": false,
        "unique": false,
        "options": {
          "collectionId": "f8w5oambwthngxo",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": null,
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
  const collection = dao.findCollectionByNameOrId("wf40hpyi2wvpb7y");

  return dao.deleteCollection(collection);
})
