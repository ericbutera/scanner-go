package main

import (
	// "storage-api/grpc"
	"flag"
	"log"
	"storage-api/config"
	"storage-api/rest"
)

var (
	path = flag.String("config_path", ".", "path to config")
	file = flag.String("config_file", "config.yaml", "config file name")
)

func main() {
	config, err := config.NewAppConfig(path, file)
	if err != nil {
		log.Fatal(err)
	}

	rest.Serve(config)
	// grpc.Serve() // TODO: move to different service
}
