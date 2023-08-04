migrate((db) => {
  const snapshot = [
    {
      "id": "3fhw2mfr9zrgodj",
      "created": "2022-12-23 22:30:35.443Z",
      "updated": "2023-05-21 11:13:12.844Z",
      "name": "hooks",
      "type": "base",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "j8mewfur",
          "name": "collection",
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
          "id": "4xcxcfuv",
          "name": "event",
          "type": "select",
          "required": true,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "values": [
              "insert",
              "update",
              "delete"
            ]
          }
        },
        {
          "system": false,
          "id": "u3bmgjpb",
          "name": "action_type",
          "type": "select",
          "required": true,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "values": [
              "command",
              "post"
            ]
          }
        },
        {
          "system": false,
          "id": "kayyu1l3",
          "name": "action",
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
          "id": "zkengev8",
          "name": "action_params",
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
          "id": "balsaeka",
          "name": "expands",
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
          "id": "emgxgcok",
          "name": "disabled",
          "type": "bool",
          "required": false,
          "unique": false,
          "options": {}
        }
      ],
      "indexes": [],
      "listRule": null,
      "viewRule": null,
      "createRule": null,
      "updateRule": null,
      "deleteRule": null,
      "options": {}
    },
    {
      "id": "_pb_users_auth_",
      "created": "2023-05-21 11:13:12.839Z",
      "updated": "2023-05-21 11:13:12.844Z",
      "name": "users",
      "type": "auth",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "users_name",
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
          "id": "users_avatar",
          "name": "avatar",
          "type": "file",
          "required": false,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "maxSize": 5242880,
            "mimeTypes": [
              "image/jpeg",
              "image/png",
              "image/svg+xml",
              "image/gif",
              "image/webp"
            ],
            "thumbs": null,
            "protected": false
          }
        }
      ],
      "indexes": [],
      "listRule": "id = @request.auth.id",
      "viewRule": "id = @request.auth.id",
      "createRule": "",
      "updateRule": "id = @request.auth.id",
      "deleteRule": "id = @request.auth.id",
      "options": {
        "allowEmailAuth": true,
        "allowOAuth2Auth": true,
        "allowUsernameAuth": true,
        "exceptEmailDomains": null,
        "manageRule": null,
        "minPasswordLength": 8,
        "onlyEmailDomains": null,
        "requireEmail": false
      }
    },
    {
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
    },
    {
      "id": "os5ld33mgj3dj7b",
      "created": "2023-05-22 05:19:19.127Z",
      "updated": "2023-05-24 08:15:52.106Z",
      "name": "repositories",
      "type": "base",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "ipzx56fk",
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
          "id": "cxcgjxz4",
          "name": "url",
          "type": "url",
          "required": true,
          "unique": false,
          "options": {
            "exceptDomains": [],
            "onlyDomains": []
          }
        },
        {
          "system": false,
          "id": "r3jb5rhz",
          "name": "status",
          "type": "select",
          "required": true,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "values": [
              "PENDING",
              "SYNCING",
              "UP-TO-DATE",
              "ERROR"
            ]
          }
        },
        {
          "system": false,
          "id": "b43e5pqo",
          "name": "repositoryCredentials",
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
      "listRule": "",
      "viewRule": "",
      "createRule": null,
      "updateRule": null,
      "deleteRule": null,
      "options": {}
    },
    {
      "id": "v81b0cv9ghqpmjr",
      "created": "2023-05-24 13:20:22.475Z",
      "updated": "2023-05-24 13:20:22.475Z",
      "name": "helmRepositoryCredentials",
      "type": "base",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "smqdrvj0",
          "name": "username",
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
          "id": "o2fcjala",
          "name": "password",
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
    },
    {
      "id": "z6e64e4y3ifjd4v",
      "created": "2023-05-24 13:20:54.976Z",
      "updated": "2023-05-24 13:20:54.976Z",
      "name": "helmRepositories",
      "type": "base",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "lltfaiwh",
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
          "id": "l1vjlw2g",
          "name": "url",
          "type": "url",
          "required": true,
          "unique": false,
          "options": {
            "exceptDomains": [],
            "onlyDomains": []
          }
        },
        {
          "system": false,
          "id": "lfh7brfk",
          "name": "helmRepositoryCredentials",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "v81b0cv9ghqpmjr",
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
    },
    {
      "id": "f8w5oambwthngxo",
      "created": "2023-05-24 13:26:55.408Z",
      "updated": "2023-05-24 14:33:55.658Z",
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
        },
        {
          "system": false,
          "id": "nyx3xebq",
          "name": "parentStage",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "f8w5oambwthngxo",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": 1,
            "displayFields": [
              "name"
            ]
          }
        },
        {
          "system": false,
          "id": "z5xx0qqt",
          "name": "deployments",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "ka4cvffqmuxczw7",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "name"
            ]
          }
        },
        {
          "system": false,
          "id": "djv5hxb6",
          "name": "template",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "9gf5p3rlwfhgh1j",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": 1,
            "displayFields": [
              "name"
            ]
          }
        },
        {
          "system": false,
          "id": "6f4m9lvy",
          "name": "keyValues",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "zp90bz3osxtcevq",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "key",
              "value"
            ]
          }
        },
        {
          "system": false,
          "id": "xirmquyd",
          "name": "secretKeyValues",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "ujpm7pjc0i3qg81",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "key"
            ]
          }
        }
      ],
      "indexes": [
        "CREATE UNIQUE INDEX `idx_Np8yfdP` ON `stages` (`keyValues`)",
        "CREATE UNIQUE INDEX `idx_skQyJ7x` ON `stages` (`secretKeyValues`)",
        "CREATE UNIQUE INDEX `idx_JXT6tA7` ON `stages` (\n  `secretKeyValues`,\n  `keyValues`\n)"
      ],
      "listRule": null,
      "viewRule": null,
      "createRule": null,
      "updateRule": null,
      "deleteRule": null,
      "options": {}
    },
    {
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
    },
    {
      "id": "ka4cvffqmuxczw7",
      "created": "2023-05-24 13:40:54.615Z",
      "updated": "2023-05-24 15:24:23.218Z",
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
          "id": "seb6k5ba",
          "name": "repositories",
          "type": "relation",
          "required": true,
          "unique": false,
          "options": {
            "collectionId": "os5ld33mgj3dj7b",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "name"
            ]
          }
        },
        {
          "system": false,
          "id": "qmvzhwkm",
          "name": "keyValues",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "zp90bz3osxtcevq",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "key",
              "value"
            ]
          }
        },
        {
          "system": false,
          "id": "vjhter9p",
          "name": "secretKeyValues",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "ujpm7pjc0i3qg81",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "key",
              "value"
            ]
          }
        }
      ],
      "indexes": [
        "CREATE UNIQUE INDEX `idx_1CGxIyJ` ON `deployments` (\n  `keyValues`,\n  `secretKeyValues`\n)",
        "CREATE UNIQUE INDEX `idx_3uvZmVx` ON `deployments` (`keyValues`)",
        "CREATE UNIQUE INDEX `idx_5PfIsEo` ON `deployments` (`secretKeyValues`)"
      ],
      "listRule": null,
      "viewRule": null,
      "createRule": null,
      "updateRule": null,
      "deleteRule": null,
      "options": {}
    },
    {
      "id": "zp90bz3osxtcevq",
      "created": "2023-05-24 14:12:20.499Z",
      "updated": "2023-05-24 14:30:43.054Z",
      "name": "keyValues",
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
          "id": "yzpsh3rf",
          "name": "displayName",
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
    },
    {
      "id": "ujpm7pjc0i3qg81",
      "created": "2023-05-24 14:19:45.983Z",
      "updated": "2023-05-24 14:30:56.612Z",
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
        },
        {
          "system": false,
          "id": "od3uiju3",
          "name": "displayName",
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
    },
    {
      "id": "9gf5p3rlwfhgh1j",
      "created": "2023-05-24 14:28:14.243Z",
      "updated": "2023-05-24 14:32:50.565Z",
      "name": "templates",
      "type": "base",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "5c1k4g8q",
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
          "id": "r5t5pphi",
          "name": "description",
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
          "id": "arhbhwoc",
          "name": "keyValues",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "zp90bz3osxtcevq",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "key",
              "displayName"
            ]
          }
        },
        {
          "system": false,
          "id": "394csawp",
          "name": "secretKeyValues",
          "type": "relation",
          "required": false,
          "unique": false,
          "options": {
            "collectionId": "ujpm7pjc0i3qg81",
            "cascadeDelete": false,
            "minSelect": null,
            "maxSelect": null,
            "displayFields": [
              "key",
              "displayName"
            ]
          }
        }
      ],
      "indexes": [
        "CREATE UNIQUE INDEX `idx_sAYJ45f` ON `templates` (`keyValues`)",
        "CREATE UNIQUE INDEX `idx_X6weBCl` ON `templates` (`secretKeyValues`)",
        "CREATE UNIQUE INDEX `idx_hNLbnBO` ON `templates` (\n  `keyValues`,\n  `secretKeyValues`\n)"
      ],
      "listRule": null,
      "viewRule": null,
      "createRule": null,
      "updateRule": null,
      "deleteRule": null,
      "options": {}
    }
  ];

  const collections = snapshot.map((item) => new Collection(item));

  return Dao(db).importCollections(collections, true, null);
}, (db) => {
  return null;
})
