package utils

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"

	"time"
)

func UploadFileToGCS(file io.Reader, fileName, bucketName string) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Failed to create GCS client: %v", err)
		return "", err
	}
	defer client.Close()

	// Generate a unique filename (optional)
	uniqueFileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), fileName)

	// Define GCS object path
	objectPath := fmt.Sprintf("uploads/%s", uniqueFileName)

	// Create object handle
	bucket := client.Bucket(bucketName)
	object := bucket.Object(objectPath)
	writer := object.NewWriter(ctx)

	// Copy file content to GCS
	if _, err := io.Copy(writer, file); err != nil {
		log.Printf("Failed to upload file to GCS: %v", err)
		return "", err
	}

	// Close writer to complete upload
	if err := writer.Close(); err != nil {
		log.Printf("Failed to finalize upload: %v", err)
		return "", err
	}

	// Return the full GCS path
	return fmt.Sprintf("gs://%s/%s", bucketName, objectPath), nil
}
