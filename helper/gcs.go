package helper

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

var (
	projectID  = os.Getenv("PROJECT_ID")
	bucketName = os.Getenv("BUCKET_NAME")
)

type ClientUploader struct {
	C1         storage.Client
	ProjectId  string
	BucketName string
	UploadPath string
}

var Uploader *ClientUploader

func init() {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		panic(err)
	}

	Uploader = &ClientUploader{
		C1:         *client,
		ProjectId:  projectID,
		BucketName: bucketName,
		UploadPath: "images/",
	}
}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	wc := c.C1.Bucket(c.BucketName).Object(c.UploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
