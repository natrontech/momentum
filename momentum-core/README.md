# Momentum-Core

The core of momentum is about managing the core structure and functionality of a core repository as described in the momentum-structure.

The core implements a REST-API empowering consumers to manage their repository instance.

# Environment variables

| Name         | Description | Mandatory
|--------------|-------------|-----------
| MOMENTUM_GIT_REPO_URL | The repository which shall be cloned on startup (if not already on disk) | yes
| MOMENTUM_GIT_USER     | Username for user who creates transactions | yes
| MOMENTUM_GIT_EMAIL    | Email for the user who creates the transactions | yes
| MOMENTUM_GIT_TOKEN    | Access token which belongs to the user who creates the transaction | yes
