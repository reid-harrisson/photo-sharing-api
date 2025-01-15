// New file for storage service
package storage

import (
	"mime/multipart"

	storage_go "github.com/supabase-community/storage-go"
)

type StorageService struct {
	client   *storage_go.Client
	bucketID string
}

func NewStorageService(client *storage_go.Client) *StorageService {
	return &StorageService{
		client:   client,
		bucketID: "image-bucket",
	}
}

func (s *StorageService) EnsureBucket() error {
	_, err := s.client.GetBucket(s.bucketID)
	if err != nil {
		_, err = s.client.CreateBucket(s.bucketID, storage_go.BucketOptions{
			Public: true,
		})
	}
	return err
}

func (s *StorageService) UploadImage(fileName string, file multipart.File) (string, error) {
	contentType := "image/jpeg"
	_, err := s.client.UploadFile(s.bucketID, fileName, file, storage_go.FileOptions{
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}

	return s.client.GetPublicUrl(s.bucketID, fileName).SignedURL, nil
}
