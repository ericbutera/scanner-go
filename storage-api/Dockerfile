FROM golang:1.18-alpine AS build

WORKDIR /src/

COPY . ./

RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
