# Momentum's Openapi generation integration

Since the generation of a good openapi spec is crucial to momentums core api, it's documented here how this was implemented.

## What is generated

- Models (folder *models*)
    - data models representing the data types accepted by momentum and sent in responses.
- API (folder *routers*)
    - api which expose momentum's core features at rest.

## Which tool is used

We use [go-swagger](https://github.com/go-swagger/go-swagger)

## Requirements

- package 'docs' (folder *docs* (this doc is in it))

## Encountered difficulties & how they were conquered (or not)



## Useful References

Stuff I was thankful to read during the integration and might help you as well:

- [perfect entrypoint](https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9)
- [generate swagger json from go code](https://goswagger.io/generate/spec.html)
- [spec on how to define what shall be generated]((https://goswagger.io/use/spec.html) )
