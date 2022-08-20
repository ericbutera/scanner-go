# Storage API

## Docs

- Generate using `make docs`
- [View](http://localhost:8080/swagger/index.html)

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

- collect raw data
- store in _unified log_ (from event streams in action)

## Telemetry

### Jaeger

- Run agent
- Enable `jaeger` in `config.yaml`
- Visit [Jaeger UI](http://localhost:16686/)

```sh
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.37
```

### DataDog

- Run agent
- Enable `data_dog` in `config.yaml`

```sh
# Run agent
docker run \
--rm \
--name datadog-agent \
-p 8126:8126 \
-v /var/run/docker.sock:/var/run/docker.sock:ro \
-v /proc/:/host/proc/:ro \
-v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
-e DD_API_KEY=$DD_API_KEY \
-e DD_APM_ENABLED=true \
-e DD_APM_NON_LOCAL_TRAFFIC=true \
-e DD_LOGS_INJECTION=true \
-e DD_PROFILING_ENDPOINT_COLLECTION_ENABLED=true \
-e DD_LOG_LEVEL=info \
datadog/agent:latest
```
