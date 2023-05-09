package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
)

type JobRepository interface {
	Create(video *domain.Job) (*domain.Job, error)
	GetOne(id string) (*domain.Job, error)
	Update(video *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{Db: db}
}

func (repo JobRepositoryDb) Create(video *domain.Job) (*domain.Job, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	err := repo.Db.Create(video).Error

	if err == nil {
		return nil, fmt.Errorf("Ocorreu um erro")
	}

	return video, err
}

func (repo JobRepositoryDb) GetOne(id string) (*domain.Job, error) {
	var video *domain.Job

	repo.Db.Preload("Video").First(&video, "id = ?", id)

	if video.ID == "" {
		return nil, fmt.Errorf("Ocorreu um erro")
	}

	return video, nil
}

func (repo JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(job).Error

	if err != nil {
		return nil, fmt.Errorf("Ocorreu um erro")
	}

	return job, nil
}
