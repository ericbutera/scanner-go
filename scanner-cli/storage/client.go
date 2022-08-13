// TODO move to external project
package client

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func NewClient() *resty.Client {
	return resty.New()
}

func Start(client *resty.Client) (string, error) {
	log.Print("started")

	resp, err := client.R().
		EnableTrace().
		Post("http://localhost:8080/scan/start")

	// TODO check status200
	// TODO retry
	// TODO auth

	data := resp.String()

	return data, err
}

func End(client *resty.Client, id string) {
	log.Print("finished")
	// POST storage-api finish
	// resp, err := client.R().
	// 	Post("http://localhost:8080/scan/{id}/finish")
}
