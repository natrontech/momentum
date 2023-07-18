# Momentum's Openapi generation integration

Since the generation of a good openapi spec is crucial to momentums core api, it's documented here how this was implemented.

We were unable to find a tool which is ready enough to generate openapi spec for version 3. Due to this we will remain on version 2 and upgrade later.

You can access the [swagger-ui](http://localhost:8080/swagger/index.html) by opening [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## What is generated

- Models (folder *models*)
    - data models representing the data types accepted by momentum and sent in responses.
- API (folder *routers*)
    - api which expose momentum's core features at rest.

## Which tool is used

We use [swag](https://github.com/swaggo/swag)

## Requirements

run `make install-swagger` to install `swag` which generates the spec.

## Hint

During the implementation the examples in of the [swag](https://github.com/swaggo/swag) repository were really helpful.
