package controllers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
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

// type Form struct {
//     Files *multipart.FileHeader `form:"files" binding:"required"`
// }

func getClient(ctx context.Context) (*storage.Client, error){
	return storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
}

func getBucket(client *storage.Client) (*storage.BucketHandle) {
	return client.Bucket(bucketName)
}


// HandleFileUploadToBucket uploads file to bucket
// you can upload multiple files or a single file.
func HandleFileUploadToBucket(c *gin.Context) {
	var err error
	ctx := appengine.NewContext(c.Request)
	storageClient, _ = getClient(ctx)
	bkt := getBucket(storageClient)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		fileName := file.Filename
		openedFile, _ := file.Open()
		uploadFileToGCS(c, &ctx, bkt, fileName, &openedFile)
    }

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
	})
}

func uploadFileToGCS(c *gin.Context, clientCtx *context.Context, bkt *storage.BucketHandle, fileName string, f *multipart.File) {
	sw := bkt.Object(fileName).NewWriter(*clientCtx)

	if _, err := io.Copy(sw, *f); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := sw.Close(); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	u, err := url.Parse("/" + bucketName + "/" + sw.Attrs().Name)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}
	fmt.Printf("%+v\n", u)
}


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