// TODO check status200
// TODO auth
// TODO content-type json
package storage

import (
	"fmt"
	"log"
	"scanner-go/retry"

	"github.com/go-resty/resty/v2"
)

type Storage struct {
	Rest    *resty.Client
	BaseUrl string `default:"http://localhost:8080"`
	ScanId  string
}

func NewStorage(baseUrl string) *Storage {
	return &Storage{
		Rest:    resty.New(),
		BaseUrl: baseUrl,
	}
}

type ScanResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func (storage *Storage) Start() error {
	// example of multiple return values from the retry function
	// err := retry.Run(func() error {
	// 	id, _err := storage._Start()
	// 	log.Printf("storage-api start id %s err %v", id, _err)
	// 	return _err
	// })
	err := retry.Run("storage-api-start", storage._Start)

	if err != nil {
		log.Printf("storage-api start error %+v", err)
	}

	return err
}

func (storage *Storage) _Start() error {
	// TODO prevent multiple calls
	data := &ScanResponse{}
	_, err := storage.Rest.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json"). // TODO there has to be a way to make this automatic
		SetResult(data).
		Post(fmt.Sprintf("%s/scan/start", storage.BaseUrl))

	if err != nil {
		return err
	}

	log.Printf("storage-api start response %+v", data)
	storage.ScanId = data.Id // <-- no return value necessary as id is set on storage

	return nil
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

func (storage *Storage) Save(data *StorageData) error {
	err := retry.Run("storage-api-save", func() error {
		return storage._Save(data)
	})

	if err != nil {
		log.Printf("storage-api start error %+v", err)
	}

	return err
}

func (storage *Storage) _Save(data *StorageData) error {
	resp, err := storage.Rest.R().
		EnableTrace().
		SetBody(data).
		Post(fmt.Sprintf("%s/scan/%s/save", storage.BaseUrl, storage.ScanId))

	log.Print("save response ", resp.String())

	return err
}

func (storage *Storage) End() error {
	return retry.Run("storage-api-end", storage._End)
}

func (storage *Storage) _End() error {
	resp, err := storage.Rest.R().
		EnableTrace().
		Post(fmt.Sprintf("%s/scan/%s/finish", storage.BaseUrl, storage.ScanId))

	log.Print("end response ", resp.String())

	return err
}
