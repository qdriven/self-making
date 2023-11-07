# Product API Spec

Use [Fern](https://docs.buildwithfern.com/):

1. Define API Spec
2. Generate API Codes
   1. Client/SDKs
   2. Server/Backend boilerplate
   3. Documentation

## Fern Overview

![](https://fern-image-hosting.s3.amazonaws.com/Fern+Overview.png)

## Installation

```sh
npm install -g fern-api
```

## Start

```sh
fern init

❯ tree fern               
fern
├── api  #your API
│   ├── definition
│   │   ├── api.yml  # API-level configuration
│   │   └── imdb.yml # endpoints, types, and errors
│   └── generators.yml
└── fern.config.json # root-level configuration

3 directories, 4 files
```

## Cheatsheet

```sh
fern generate
fern check
fern add 
fern init
fern write definition # covert openapi to fern definition
fern upgrade
fern register
fern format
fern help
fern version

```

## Why

***schema-first automation for API development:***
- Schema-first API design
  - Best practices built-in
  - A single source of truth
  - A simple format encourages collaboration
- Product velocity
  - Think in terms of operations, not HTTP
  - Type safety
  - Auto-completion in code editors

## OpenAPI Way Example 

> # 📑 Qase API Specification
> 
> Based on [OpenAPI](https://en.wikipedia.org/wiki/OpenAPI_Specification) specifications.

## Useful documentations

- [OpenAPI Initiative](https://www.openapis.org).
- [OpenAPI Specification](https://swagger.io/specification/).
- [Swagger Documentation](https://swagger.io/docs/).

## Useful resources

- [OpenAPI Tools](https://github.com/OpenAPITools).
  - [openapi-generator-cli](https://github.com/OpenAPITools/openapi-generator-cli).
- [Swagger at GitHub](https://github.com/swagger-api).
  - [Swagger Codegen](https://github.com/swagger-api/swagger-codegen).
  - [Swagger UI](https://github.com/swagger-api/swagger-ui).
- [ReadMe at GitHub](https://github.com/readmeio).
  - [rdme](https://github.com/readmeio/rdme).
- [Redocly](https://github.com/Redocly).
  - [ReDoc](https://github.com/Redocly/redoc).
- [API Dev Tools](https://github.com/APIDevTools).
  - [swagger-cli](https://github.com/APIDevTools/swagger-cli).
- Templates
  - [OpenAPI Boilerplate](https://github.com/dgarcia360/openapi-boilerplate).
  - [OpenAPI Definition Starter](https://github.com/Redocly/openapi-starter).

## Useful articles

- [How to split a large OpenAPI document into multiple files](https://davidgarcia.dev/posts/how-to-split-open-api-spec-into-multiple-files/).
do



## References


- [qase-spec](https://github.com/qase-tms/specs.git)



