package service

import (
	"context"
	"mime/multipart"
	"restapi/configs"
	"restapi/internal/entity"
	"restapi/pkg/repository"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadImageServise struct {
	repo repository.UploadImage
}

func NewUploadImageService(repo repository.UploadImage) *UploadImageServise {
	return &UploadImageServise{repo: repo}
}

//func (s *UploadImageServise) Upload(userId int, image entity.Image) (int, error) {
//	image.Image = repository.UploadToCloudinary
//}

func (s *UploadImageServise) Upload(userId int, image entity.Image, file multipart.File, filePath string) (int, error) {
	ctx := context.Background()
	cld, err := configs.SetupCloudinary()
	if err != nil {
		return 0, err
	}

	// create upload params
	uploadParams := uploader.UploadParams{
		PublicID:     filePath,
		ResourceType: "image",
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return 0, err
	}

	image.Image = result.SecureURL
	return s.repo.UploadImage(userId, image)
}

func (s *UploadImageServise) GetAll(userId int) ([]entity.Image, error) {
	return s.repo.GetAll(userId)
}

func (s *UploadImageServise) GetById(userId, imageId int) (entity.Image, error) {
	return s.repo.GetById(userId, imageId)
}

func (s *UploadImageServise) Delete(userId, imageId int) error {
	return s.repo.Delete(userId, imageId)
}
