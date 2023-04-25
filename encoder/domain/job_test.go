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

	job, err := domain.NewJob(output, status, &video)

	require.NotNil(t, job)
	require.Nil(t, err)
}

func TestJobIdIsNotUuid(t *testing.T) {
	output := ""
	status := ""
	video := domain.Video{}

	job, err := domain.NewJob(output, status, &video)

	require.Nil(t, job)
	require.Error(t, err)
}
