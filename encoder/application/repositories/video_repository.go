package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
)

type VideoRepository interface {
	Create(video *domain.Video) (*domain.Video, error)
	GetOne(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{Db: db}
}

func (repo VideoRepositoryDb) Create(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	err := repo.Db.Create(video).Error

	if err == nil {
		return nil, fmt.Errorf("Ocorreu um erro")
	}

	return video, err
}

func (repo VideoRepositoryDb) GetOne(id string) (*domain.Video, error) {
	var video *domain.Video

	repo.Db.Preload("Jobs").First(&video, "id = ?", id)

	if video.ID == "" {
		return nil, fmt.Errorf("Ocorreu um erro")
	}

	return video, nil
}
