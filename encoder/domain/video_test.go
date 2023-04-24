package domain_test

import (
	"encoder/domain"
	"testing"

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
