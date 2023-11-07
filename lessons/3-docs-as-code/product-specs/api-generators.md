# Fern as API builder

>
Google built gRPC. Amazon built Smithy. Facebook built GraphQL. Palantir built Conjure.

What's good:
1. built-in
2. a single source of truth
3. simple format for collaborations

***schema-first*** automation for API development: 

basic workflow:
1. define api
2. use fern to generate codes
   1. server sides
   2. client sides
   3. postman collection
   4. openapi spec

## install fern-api

```sh
npm install -g fern-api
```

## Write about spec

```sh
fern init

❯ tree fern              
fern
├── api
│   ├── definition
│   │   ├── api.yml
│   │   └── imdb.yml
│   └── generators.yml
└── fern.config.json

3 directories, 4 files
```

## start from openapi

```sh
fern init --openapi <file>
```

## Generate Code

```sh
fern generate
```






