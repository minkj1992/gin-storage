package services

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)


const (
	bucketName = "voda-storage"
)

var (
	storageClient *storage.Client
)

func getClient(ctx context.Context) (*storage.Client, error){
	return storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
}

func getBucket(client *storage.Client) (*storage.BucketHandle) {
	return client.Bucket(bucketName)
}


// HandleFileUploadToBucket uploads file to bucket
func HandleFileUploadToBucket(c *gin.Context) {
	var err error

	ctx := appengine.NewContext(c.Request)
	storageClient, err = getClient(ctx)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	defer f.Close()
	bkt := getBucket(storageClient)
	sw := bkt.Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	u, err := url.Parse("/" + bucketName + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"pathname": u.EscapedPath(),
	})
}

// streaming?
func HandleFileDownloadFromBucket(c *gin.Context) {
	var err error
	tmpFileName := "animal.jpeg"
	ctx := appengine.NewContext(c.Request)
	storageClient, err = getClient(ctx)
	bkt := getBucket(storageClient)
	obj := bkt.Object(tmpFileName)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	// Read it back.
	r, err := obj.NewReader(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer r.Close()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		// TODO: Handle error.
	}
}