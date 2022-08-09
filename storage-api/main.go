package main

import (
	"storage-api/grpc"
	"storage-api/rest"
)

func main() {
	grpc.Serve()
	rest.Serve()
}
