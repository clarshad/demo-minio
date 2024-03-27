package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "localhost:9000"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	ctx := context.Background()

	// create a minio client
	c, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// create a bucket
	bucketName := "demo-bucket"
	c.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// check if bucket already exist
		exists, errBucketExists := c.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("successfully created bucket %v", bucketName)
	}

	// upload a test file
	info, err := c.FPutObject(ctx, bucketName, "mytestfile", "./testfile", minio.PutObjectOptions{ContentType: "application/json"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("successfully uploaded file of size %v", info.Size)
}
