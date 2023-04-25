package domain_test

import (
	"encoder/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidationJob(t *testing.T) {
	output := "."
	status := "processing"
	video := domain.Video{}

	_, err := domain.NewJob(output, status, &video)

	require.Nil(t, err)
}

// func TestJobIdIsNotUuid(t *testing.T) {
// 	job := domain.NewJob()

// 	job.ID = "82121021"

// 	err := job.Validate()

// 	require.Error(t, err)
// }

// func TestJobValidation(t *testing.T) {
// 	job := domain.NewJob()

// 	video := domain.NewVideo()

// 	job.ID = uuid.NewV4().String()
// 	job.OutputBucketPath = "."
// 	job.Status = "success"
// 	job.Video = video
// 	job.Error = "not"
// 	job.CreatedAt = time.Now()

// 	err := job.Validate()

// 	require.Nil(t, err)
// }
