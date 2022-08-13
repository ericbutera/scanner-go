// TODO move to external project
package storage

import (
	"log"

	"github.com/go-resty/resty/v2"
)

type Storage struct {
	Rest   *resty.Client
	ScanId string
}

func NewStorage() *Storage { // *resty.Client {
	return &Storage{
		Rest: resty.New(),
	}
	//return resty.New()
}

func (storage *Storage) Start() (string, error) {
	log.Print("started")

	resp, err := storage.Rest.R().
		EnableTrace().
		Post("http://localhost:8080/scan/start")

	// TODO check status200
	// TODO retry
	// TODO auth

	data := resp.String()

	return data, err
}

type StorageData struct {
	ScanId string
	Data   any
}

func (storage *Storage) NewStorageData(data any) *StorageData {
	return &StorageData{
		ScanId: storage.ScanId,
		Data:   data,
	}
}

func (storage *Storage) Save(data *StorageData) (string, error) {
	log.Print("saved")
	resp, err := storage.Rest.R().
		EnableTrace().
		SetBody(data).
		Post("http://localhost:8080/scan/{id}/save")

	body := resp.String()
	return body, err
}

func (storage *Storage) End() (data string, err error) {
	log.Print("finished")
	// POST storage-api finish
	resp, err := storage.Rest.R().
		EnableTrace().
		Post("http://localhost:8080/scan/{id}/finish")

	data = resp.String()
	return data, err
}
