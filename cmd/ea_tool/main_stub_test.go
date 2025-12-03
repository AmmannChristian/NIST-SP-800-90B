//go:build teststub

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunCLI_StdinSuccessWithStub(t *testing.T) {
	var out bytes.Buffer
	data := []byte{1, 2, 3, 4}
	code := runCLI([]string{"-non-iid", "-bits", "8"}, bytes.NewReader(data), &out, &out)
	assert.Equal(t, 0, code)
	assert.Contains(t, out.String(), "Entropy Assessment Results")
}

func TestRunCLI_OutputFileSuccess(t *testing.T) {
	var out bytes.Buffer
	data := []byte{1, 2, 3, 4}
	tmpFile := filepath.Join(t.TempDir(), "result.json")

	code := runCLI([]string{"-non-iid", "-bits", "8", "-output", tmpFile}, bytes.NewReader(data), &out, &out)
	require.Equal(t, 0, code)
	assert.Contains(t, out.String(), "Results written to")

	raw, err := os.ReadFile(tmpFile)
	require.NoError(t, err)

	var got JSONOutput
	require.NoError(t, json.Unmarshal(raw, &got))
	assert.Equal(t, "Non-IID", got.TestType)
	assert.Equal(t, 0, got.ErrorCode)
	assert.Equal(t, len(data), got.DataSize)
}

func TestRunCLI_IIDModeSuccess(t *testing.T) {
	var out bytes.Buffer
	data := []byte{1, 2, 3, 4}

	code := runCLI([]string{"-iid", "-bits", "8"}, bytes.NewReader(data), &out, &out)
	require.Equal(t, 0, code)
	assert.Contains(t, out.String(), "Test Type:       IID")
	assert.Contains(t, out.String(), "H_bitstring")
}
