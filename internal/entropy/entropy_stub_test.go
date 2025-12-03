//go:build teststub

package entropy

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The following success tests rely on the teststub build tag to avoid CGO calls.
func TestAssessIID_SuccessStub(t *testing.T) {
	assessment := NewAssessment()
	assessment.SetVerbose(0)

	res, err := assessment.AssessIID([]byte{1, 2, 3, 4}, 8)
	require.NoError(t, err)
	assert.Equal(t, 7.5, res.MinEntropy)
	assert.Equal(t, IID, res.TestType)
}

func TestAssessNonIID_SuccessStub(t *testing.T) {
	assessment := NewAssessment()
	assessment.SetVerbose(0)

	res, err := assessment.AssessNonIID([]byte{1, 2, 3, 4}, 8)
	require.NoError(t, err)
	assert.Equal(t, 6.5, res.MinEntropy)
	assert.Equal(t, NonIID, res.TestType)
}

func TestAssessFile_SuccessStub(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "data.bin")
	require.NoError(t, os.WriteFile(file, []byte{1, 2, 3, 4}, 0o644))

	assessment := NewAssessment()
	assessment.SetVerbose(0)

	res, err := assessment.AssessFile(file, 8, IID)
	require.NoError(t, err)
	assert.Equal(t, 7.5, res.MinEntropy)
	assert.Equal(t, IID, res.TestType)
}
