package handlers

import (
	"fmt"
	"gin-test/responses"
	"gin-test/server"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"

	"github.com/gin-gonic/gin"
)

type HandlerUpload struct {
	Server *server.Server
}

func NewHandlerUpload(server *server.Server) *HandlerUpload {
	return &HandlerUpload{
		Server: server,
	}
}

// UploadImage godoc
// @Summary Upload image to Supabase
// @Schemes
// @Description Upload image to Supabase
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file to upload"
// @Success 200 {object} responses.Data
// @Router /upload [post]
func (handler *HandlerUpload) UploadImage(context *gin.Context) {
	// Get the file from the request
	file, err := context.FormFile("image")
	if err != nil {
		responses.ErrorResponse(context, http.StatusBadRequest, "No file uploaded")
		return
	}

	// Validate file type
	if !isValidImageType(file) {
		responses.ErrorResponse(context, http.StatusBadRequest, "Invalid file type. Only images are allowed")
		return
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		responses.ErrorResponse(context, http.StatusInternalServerError, "Error processing file")
		return
	}
	defer fileContent.Close()

	// Upload to Supabase Storage
	bucketId := "image-bucket"

	_, err = handler.Server.SupabaseClient.GetBucket(bucketId)
	if err != nil {
		_, err = handler.Server.SupabaseClient.CreateBucket(bucketId, storage_go.BucketOptions{
			Public: true,
		})

		if err != nil {
			responses.ErrorResponse(context, http.StatusInternalServerError, "Error creating Supabase Bucket: "+err.Error())
			return
		}
	}

	contextType := context.ContentType()
	_, err = handler.Server.SupabaseClient.UploadFile(bucketId, newFileName, fileContent, storage_go.FileOptions{
		ContentType: &contextType,
	})
	if err != nil {
		responses.ErrorResponse(context, http.StatusInternalServerError, "Error uploading to Supabase: "+err.Error())
		return
	}

	// Get the public URL
	publicURL := handler.Server.SupabaseClient.GetPublicUrl(bucketId, newFileName)

	responses.Response(context, http.StatusOK, gin.H{
		"message": "Image uploaded successfully!",
		"url":     publicURL,
	})
}

func isValidImageType(file *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}
	return allowedTypes[file.Header.Get("Content-Type")]
}
