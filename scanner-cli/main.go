package main

import (
	"log"
	"scanner-go/config"
	storage "scanner-go/storage"
	// "scanner-go/provider/aws"
	// "scanner-go/provider/azure"
	// "scanner-go/provider/gcp"
)

func main() {
	config_path := "."
	conf, err := config.NewAppConfig(config_path)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config: %+v", conf)

	client := storage.NewClient()

	scan_id, err := storage.Start(client)
	if err != nil {
		log.Fatal(err)
	}

	// gcp.Scan(conf)
	// aws.Scan(conf)
	// azure.Scan(conf)

	storage.End(client, scan_id)
}
