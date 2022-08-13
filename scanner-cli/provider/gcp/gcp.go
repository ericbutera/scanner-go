// https://console.cloud.google.com
package gcp

import (
	"context"
	"fmt"
	"log"
	"scanner-go/config"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Scan(conf config.AppConfig) {
	creds := option.WithCredentialsFile(conf.GCPServiceAccount)
	fmt.Printf("creds %+v", creds)
	GetGCS(conf, creds)
}

func GetGCS(conf config.AppConfig, creds option.ClientOption) ([]string, error) {
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
	}
	return buckets, nil
}
