// https://console.cloud.google.com
package gcp

import (
	"context"
	"fmt"
	"log"
	"scanner-go/config"
	_storage "scanner-go/storage"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Scan(conf config.AppConfig, store *_storage.Storage) {
	creds := option.WithCredentialsFile(conf.GCPServiceAccount)
	fmt.Printf("creds %+v", creds)
	GetGCS(conf, creds, store)
}

func GetGCS(conf config.AppConfig, creds option.ClientOption, store *_storage.Storage) ([]string, error) {
	projectID := conf.GCPProjectId

	// us-east1
	// scanner-go-gcp
	// https://console.cloud.google.com/storage/browser/scanner-go-gcp
	// gs://scanner-go-gcp
	ctx := context.Background()
	client, err := storage.NewClient(ctx, creds)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var buckets []string
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
		fmt.Printf("Bucket: %v\n", battrs.Name)
		//data := struct{ BucketName string }{BucketName: battrs.Name}
		//bucket := struct{ BucketName string }{BucketName: battrs.Name}
		bucket := map[string]interface{}{"BucketName": battrs.Name}
		store.Save(&_storage.StorageData{
			ScanId: "asdf",
			Data:   bucket,
		})
	}
	return buckets, nil
}
