package handlers

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"photo-sharing-api/responses"
	"photo-sharing-api/server"
	"photo-sharing-api/services/storage"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	Server  *server.Server
	Service *storage.StorageService
}

func NewStorageHandler(server *server.Server) *StorageHandler {
	return &StorageHandler{
		Server:  server,
		Service: storage.NewStorageService(server.SupabaseClient),
	}
}

// Storage godoc
// @Summary Upload image to Supabase Storage
// @Schemes
// @Description Upload image to Supabase Storage
// @Tags Storage
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file to upload"
// @Success 200 {object} responses.ResponseStorage
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /storage [post]
func (handler *StorageHandler) UploadImage(context *gin.Context) {
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
	newFileName := uuid.New().String() + filepath.Ext(file.Filename)

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		responses.ErrorResponse(context, http.StatusInternalServerError, "Error processing file")
		return
	}
	defer fileContent.Close()

	err = handler.Service.EnsureBucket()
	if err != nil {
		responses.ErrorResponse(context, http.StatusInternalServerError, "Error creating Supabase Bucket")
		return
	}

	// Upload and get the public URL
	publicURL, err := handler.Service.UploadImage(newFileName, fileContent)
	if err != nil {
		responses.ErrorResponse(context, http.StatusInternalServerError, "Error uploading to Supabase")
		return
	}

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
