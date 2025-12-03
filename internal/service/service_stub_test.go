//go:build teststub

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Success paths rely on the teststub build tag to avoid CGO.
func TestService_AssessIID_SuccessStub(t *testing.T) {
	svc := NewService()
	res, err := svc.AssessIID([]byte{1, 2, 3, 4}, 8)
	require.NoError(t, err)
	assert.Equal(t, 7.5, res.MinEntropy)
}

func TestService_AssessNonIID_SuccessStub(t *testing.T) {
	svc := NewService()
	res, err := svc.AssessNonIID([]byte{1, 2, 3, 4}, 8)
	require.NoError(t, err)
	assert.Equal(t, 6.5, res.MinEntropy)
}

func TestService_AssessIID_AssessmentError(t *testing.T) {
	svc := NewService()

	_, err := svc.AssessIID([]byte{0xFF, 1, 2, 3}, 8)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "IID assessment failed")
}

func TestService_AssessNonIID_AssessmentError(t *testing.T) {
	svc := NewService()

	_, err := svc.AssessNonIID([]byte{0xFF, 1, 2, 3}, 8)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Non-IID assessment failed")
}
