package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidationIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestValidationIfVideoIdIsNotUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "12345"

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "82121021"

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceId = "a"
	video.FilePath = "."
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
