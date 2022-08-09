// https://aws.amazon.com/
// https://github.com/aws/aws-sdk-go-v2/tree/main/aws
package aws

import (
	"context"
	"encoding/json"
	"log"
	appconfig "scanner-go/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Scan(app_config appconfig.AppConfig) {
	aws_cfg, err := NewCredsFromAppConfig(app_config)
	if err != nil {
		log.Fatal(err)
	}

	GetS3(aws_cfg)
	// GetS3(sess)
}

func NewCredsFromAppConfig(app_config appconfig.AppConfig) (aws.Config, error) {
	// region := config.WithRegion("us-east-1")
	region := config.WithRegion(app_config.Region)

	// Load the Shared AWS Configuration (~/.aws/config)
	aws_cfg, err := config.LoadDefaultConfig(context.TODO(), region)
	return aws_cfg, err
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

func GetS3(cfg aws.Config) {
	// Create an Amazon S3 service client
	// Get the first page of results for ListObjectsV2 for a bucket
	client := s3.NewFromConfig(cfg)
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("aws-scanner-go"),
	})
	if err != nil {
		log.Fatal(err)
	}

	SaveResponse(output)
	// {"CommonPrefixes":null,"Contents":[{"ChecksumAlgorithm":null,"ETag":"\"135..."","Key":"file-name.txt","LastModified":"2022-07-31T23:42:34Z","Owner":null,"Size":19,"StorageClass":"STANDARD"}],"ContinuationToken":null,"Delimiter":null,"EncodingType":"","IsTruncated":false,"KeyCount":1,"MaxKeys":1000,"Name":"bucket-name","NextContinuationToken":null,"Prefix":"","StartAfter":null,"ResultMetadata":{}}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}

func SaveResponse(v any) {
	// Send to storage-api
	// TODO save any struct to data lake for later analysis
	json, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(json))
}
