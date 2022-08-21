# Storage API

This is a project I am using to learn about cloud orchestration.

## Docs

- Generate using `make docs`
- [View](http://localhost:8080/swagger/index.html)

## Goals

- Ability to replay processing of raw data
- fast storage of raw scanner data (this might happen in a different [project](https://github.com/ericbutera/airflow-instance) now.)
- /docs use <https://github.com/swaggo/gin-swagger>
- gRPC communication (premature optimization)
  - [gin gRPC example](https://github.com/gin-gonic/examples/blob/master/grpc/gin/main.go)
- telemetry and observability
  - see what is available in jaeger and how it compares to DataDog
  - see what can be orchestrated within a k8s cluster (service mesh offers this for free)
- helm chart
  - configuration in k8s?

## Development Log

- use [gin](https://github.com/gin-gonic/gin)
  - (claims to support gRPC)
- init

## Concept

- collect raw data
- store in _unified log_ (from event streams in action)
- deploy with kubernetes

## Kubernetes

One of the interesting issues I ran into here was container architecture. Docker has created a `buildx` [plugin](https://docs.docker.com/build/buildx/) to help support multi-architecture build support. When I created the image on my M1 (arm64) it didn't work on my linux k8s host (amd64).

### Deployment

I created some scattered specs in the kubernetes directory. These are enough to deploy the storage-api app inside k8s. Next step will be to create a helm chart.

### Telepresence

A major goal for this project is to be able to debug the microservice on my local machine instead of deploying to the cluster to see changes. Telepresence is an easy way to accomplish this.

- [telepresence intercepts](https://www.telepresence.io/docs/v2.0/howtos/intercepts/)

```sh
# microk8s running @ https://storage-api.k8s
# telepresence will route traffic to laptop instead of k8s service:
telepresence intercept storage-api --port 8080 --env-file ~/tmp/storage-api-intercept.env
# run service locally
make run # or use vscode debugger for live debugging!
```

Note: the intercepts became stuck and a `killall telepresence` fixed the issue.

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
