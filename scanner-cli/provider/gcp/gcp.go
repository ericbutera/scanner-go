// https://console.cloud.google.com
package gcp

import (
	"context"
	"fmt"
	"log"
	"scanner-go/config"
	_storage "scanner-go/storage" // TODO fix name
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Scan(conf config.AppConfig, store *_storage.Storage) error {
	log.Print("Scan GCP")

	creds := option.WithCredentialsFile(conf.GCPServiceAccount)

	fmt.Printf("creds %+v", creds)

	// individual services will not stop the scan
	if err := GCS(conf, creds, store); err != nil {
		log.Println(err)
	}

	return nil
}

// Google Cloud Storage
func GCS(conf config.AppConfig, creds option.ClientOption, store *_storage.Storage) error {
	// it, err := GetGCS(conf, creds, store) // TODO
	ctx := context.Background()
	client, err := storage.NewClient(ctx, creds)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	it := client.Buckets(ctx, conf.GCPProjectId)
	if err != nil {
		return err
	}

	for {
		bucket, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("Bucket: %v\n", bucket.Name)
		SaveGCS(store, bucket.Name)
	}

	return nil
}

// TODO: why doesnt this work? (panic: runtime error: invalid memory address or nil pointer dereference)
// Fetch buckets
// func GetGCS(conf config.AppConfig, creds option.ClientOption, store *_storage.Storage) (*storage.BucketIterator, error) {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx, creds)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()
// 	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
// 	defer cancel()
// 	it := client.Buckets(ctx, conf.GCPProjectId)
// 	return it, nil
// }

type GCSBucket struct {
	Name string
}

// Save buckets
func SaveGCS(store *_storage.Storage, bucket string) {
	data := store.NewStorageData(&GCSBucket{
		Name: bucket,
	})
	store.Save(data)
}
