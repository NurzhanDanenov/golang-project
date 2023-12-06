package service

import (
	"context"
	"mime/multipart"
	"restapi/internal/entity"
	"restapi/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mocks_test.go

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	Users(ctx context.Context) ([]*entity.User, error)
}

type UploadImage interface {
	Upload(userId int, image entity.Image, file multipart.File, filePath string) (int, error)
	GetAll(userId int) ([]entity.Image, error)
	GetById(userId, imageId int) (entity.Image, error)
	Delete(userId, imageId int) error
}

type Service struct {
	Authorization
	UploadImage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UploadImage:   NewUploadImageService(repos.UploadImage),
	}
}
