# Storage API

## Goals

- Ability to replay processing of raw data
- fast storage of raw scanner data
- /docs use <https://github.com/swaggo/gin-swagger>
- gRPC communication (premature optimization)
  - [gin gRPC example](https://github.com/gin-gonic/examples/blob/master/grpc/gin/main.go)

## Development Log

- use [gin](https://github.com/gin-gonic/gin)
  - (claims to support gRPC)
- init

## Concept

```
1. Await CLI submitted scanner data
2. Save raw scanner data
3. Enqueue worker to process raw data
  - this is where things get interesting
  - storage api will have to process the raw data without schema or version info
```
