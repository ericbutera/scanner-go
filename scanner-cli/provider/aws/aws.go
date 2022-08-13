// https://aws.amazon.com/
// https://github.com/aws/aws-sdk-go-v2/tree/main/aws
// TODO: https://github.com/DataDog/dd-trace-go/blob/v1.40.1/contrib/aws/aws-sdk-go/aws/example_test.go
package aws

import (
	"context"
	"log"

	appconfig "scanner-go/config"
	_storage "scanner-go/storage" // TODO fix name

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Scan(app_config appconfig.AppConfig, store *_storage.Storage) error {
	aws_cfg, err := NewCredsFromAppConfig(app_config)
	if err != nil {
		log.Printf("AWS credential error: %s", err)
		return err
	}

	// individual services should not stop the scan
	if err = GetS3(aws_cfg, store); err != nil {
		log.Print("S3 error ", err)
	}

	return nil
}

func NewCredsFromAppConfig(app_config appconfig.AppConfig) (aws.Config, error) {
	// region := config.WithRegion("us-east-1")
	region := config.WithRegion(app_config.Region)

	// Load the Shared AWS Configuration (~/.aws/config)
	aws_cfg, err := config.LoadDefaultConfig(context.TODO(), region)
	return aws_cfg, err
}

func GetS3(cfg aws.Config, store *_storage.Storage) error {
	// Create an Amazon S3 service client
	// Get the first page of results for ListObjectsV2 for a bucket
	client := s3.NewFromConfig(cfg)
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("aws-scanner-go"),
	})
	if err != nil {
		return err
	}

	// output={"CommonPrefixes":null,"Contents":[{"ChecksumAlgorithm":null,"ETag":"\"135..."","Key":"file-name.txt","LastModified":"2022-07-31T23:42:34Z","Owner":null,"Size":19,"StorageClass":"STANDARD"}],"ContinuationToken":null,"Delimiter":null,"EncodingType":"","IsTruncated":false,"KeyCount":1,"MaxKeys":1000,"Name":"bucket-name","NextContinuationToken":null,"Prefix":"","StartAfter":null,"ResultMetadata":{}}

	log.Println("first page results:")
	// for _, object := range output.Contents {
	// 	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	// }

	store.Save(store.NewStorageData(output))

	return nil
}

// TODO
// func NewCredsFromFile(app_config appconfig.AppConfig) (session.Session, error) {
// 	region := "us-east-1" // aws.String("us-east-1")
// 	creds := &credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN")
// 	aws_config := &aws.Config{
// 		Region:      region,
// 		Credentials: creds,
// 	}
// 	sess, err := session.NewSession(aws_config)
// 	return *sess, err
// }
