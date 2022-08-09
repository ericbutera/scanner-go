# Scanner Go

## Goals

- Scanning process is so simple it really cannot fail
- Concurrency by sub-process (maybe worker pool if confidence exceptions wont crash the thing)
- Save raw responses
  - [AVRO???](https://github.com/linkedin/goavro)
  - [msgpack](https://github.com/vmihailenco/msgpack)
- Small binary
  - small attack surface
  - quick building
- Client side rate limiting
- Metrics & logs
  - run duration
  - scanner totals

## Development Log

- TODO: create storage api in another project
- implement azure bucket
- implement gcp bucket
- implement aws bucket
- full docker image ~300mb
- from scratch 8.8mb

## Research

- [mockery](https://github.com/jaytaylor/mockery-example/blob/master/main.go)
- [godotenv](https://github.com/joho/godotenv)

## Concept

```
1. CLI scan infra
2. CLI emits findings to StorageAPI
```
