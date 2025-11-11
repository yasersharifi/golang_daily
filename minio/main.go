package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// ---------------------------------------------------------------------
// 1. Client creation
// ---------------------------------------------------------------------
func newClient(endpoint, accessKey, secretKey string, secure bool) (*minio.Client, error) {
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
}

// ---------------------------------------------------------------------
// 2. Bucket
// ---------------------------------------------------------------------
func ensureBucket(c *minio.Client, name, region string) {
	ctx := context.Background()
	if err := c.MakeBucket(ctx, name, minio.MakeBucketOptions{Region: region}); err != nil {
		exists, e2 := c.BucketExists(ctx, name)
		if e2 != nil || !exists {
			log.Fatalf("MakeBucket error: %v", err)
		}
		fmt.Printf("Bucket %q already exists\n", name)
		return
	}
	fmt.Printf("Bucket %q created\n", name)
}

// ---------------------------------------------------------------------
// 3. Upload (image / video)
// ---------------------------------------------------------------------
func upload(c *minio.Client, bucket, object, localPath, contentType string) {
	ctx := context.Background()
	info, err := c.FPutObject(ctx, bucket, object, localPath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalf("Upload %s failed: %v", object, err)
	}
	fmt.Printf("Uploaded %s (%d bytes, ETag %s)\n", object, info.Size, info.ETag)
}

// ---------------------------------------------------------------------
// 4. Presigned download URL
// ---------------------------------------------------------------------
func presignGet(c *minio.Client, bucket, object string, expiry time.Duration) string {
	ctx := context.Background()
	u, err := c.PresignedGetObject(ctx, bucket, object, expiry, url.Values{})
	if err != nil {
		log.Fatalf("Presign failed: %v", err)
	}
	return u.String()
}

// ---------------------------------------------------------------------
// 5. List objects
// ---------------------------------------------------------------------
func listObjects(c *minio.Client, bucket string) {
	ctx := context.Background()
	for obj := range c.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: true}) {
		if obj.Err != nil {
			log.Printf("list err: %v", obj.Err)
			continue
		}
		fmt.Printf("  • %s  (%d B)\n", obj.Key, obj.Size)
	}
}

// ---------------------------------------------------------------------
// 6. (Optional) Public-read policy
// ---------------------------------------------------------------------
func setPublicRead(c *minio.Client, bucket string) {
	ctx := context.Background()
	policy := fmt.Sprintf(`{
		"Version":"2012-10-17",
		"Statement":[{
			"Effect":"Allow",
			"Principal":"*",
			"Action":"s3:GetObject",
			"Resource":"arn:aws:s3:::%s/*"
		}]
	}`, bucket)

	if err := c.SetBucketPolicy(ctx, bucket, policy); err != nil {
		log.Fatalf("Set policy failed: %v", err)
	}
	fmt.Printf("Bucket %q is now public-read\n", bucket)
}

// ---------------------------------------------------------------------
// 5. List objects
// ---------------------------------------------------------------------
func listObjects(c *minio.Client, bucket string) {
	ctx := context.Background()
	for obj := range c.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: true}) {
		if obj.Err != nil {
			log.Printf("list err: %v", obj.Err)
			continue
		}
		fmt.Printf("  • %s  (%d B)\n", obj.Key, obj.Size)
	}
}

// ---------------------------------------------------------------------
// 6. (Optional) Public-read policy
// ---------------------------------------------------------------------
func setPublicRead(c *minio.Client, bucket string) {
	ctx := context.Background()
	policy := fmt.Sprintf(`{
		"Version":"2012-10-17",
		"Statement":[{
			"Effect":"Allow",
			"Principal":"*",
			"Action":"s3:GetObject",
			"Resource":"arn:aws:s3:::%s/*"
		}]
	}`, bucket)

	if err := c.SetBucketPolicy(ctx, bucket, policy); err != nil {
		log.Fatalf("Set policy failed: %v", err)
	}
	fmt.Printf("Bucket %q is now public-read\n", bucket)
}

// ---------------------------------------------------------------------
// 7. NEW: Delete object
// ---------------------------------------------------------------------
func deleteObject(c *minio.Client, bucket, object string) {
	ctx := context.Background()
	opts := minio.RemoveObjectOptions{
		// Optional: specify version ID for versioned buckets
		// VersionID: "abc123",
	}

	err := c.RemoveObject(ctx, bucket, object, opts)
	if err != nil {
		log.Fatalf("Failed to delete object %s: %v", object, err)
	}
	fmt.Printf("Deleted object: %s\n", object)
}

// ---------------------------------------------------------------------
// main
// ---------------------------------------------------------------------
func main() {
	// ---- configuration ------------------------------------------------
	const (
		endpoint        = "localhost:9000" // change if MinIO runs elsewhere
		accessKey       = "minioadmin"
		secretKey       = "minioadmin"
		useSSL          = false
		bucketName      = "demo-bucket"
		region          = "us-east-1"
		imageLocalPath  = "./sample.jpg"
		videoLocalPath  = "./sample.mp4"
		imageObjectName = "photos/sample.jpg"
		videoObjectName = "videos/sample.mp4"
	)

	// ---- client -------------------------------------------------------
	client, err := newClient(endpoint, accessKey, secretKey, useSSL)
	if err != nil {
		log.Fatalf("client init: %v", err)
	}
	fmt.Println("Connected to MinIO")

	// ---- bucket -------------------------------------------------------
	ensureBucket(client, bucketName, region)

	// ---- upload -------------------------------------------------------
	if _, err := os.Stat(imageLocalPath); err == nil {
		upload(client, bucketName, imageObjectName, imageLocalPath, "image/jpeg")
	} else {
		fmt.Println("image file missing, skipping upload")
	}
	if _, err := os.Stat(videoLocalPath); err == nil {
		upload(client, bucketName, videoObjectName, videoLocalPath, "video/mp4")
	} else {
		fmt.Println("video file missing, skipping upload")
	}

	// ---- presigned URL ------------------------------------------------
	if _, err := os.Stat(imageLocalPath); err == nil {
		url := presignGet(client, bucketName, imageObjectName, 24*time.Hour)
		fmt.Printf("\nDownload link (image, 24h):\n%s\n", url)
	}

	// ---- list ---------------------------------------------------------
	fmt.Println("\nObjects in bucket:")
	listObjects(client, bucketName)

	// ---- (optional) make bucket public --------------------------------
	// setPublicRead(client, bucketName)

	// ---- DELETE OBJECT ------------------------------------------------
	// Let's delete the uploaded image
	deleteObject(client, bucketName, imageObjectName)

	// ---- list again to confirm deletion -------------------------------
	fmt.Println("\nObjects in bucket (after delete):")
	listObjects(client, bucketName)

	// ---- (optional) make bucket public --------------------------------
	// setPublicRead(client, bucketName)
}