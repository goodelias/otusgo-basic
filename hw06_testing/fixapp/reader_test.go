package fixapp

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadJSON(t *testing.T) {
	testData := []Employee{
		{1, 23, "Ilya", 11},
		{2, 30, "Alec", 12},
	}

	// Create temp file
	tempFile, err := os.CreateTemp("", "data.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	// Delete temp file after testing
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("failed to remove temp file: %v", err)
		}
	}(tempFile.Name())

	// Writing test data to a file
	if err := json.NewEncoder(tempFile).Encode(testData); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	testCases := []struct {
		name     string
		filePath string
		want     []Employee
		wantErr  bool
	}{
		{
			name:     "valid JSON",
			filePath: tempFile.Name(),
			want:     testData,
			wantErr:  false,
		},
		{
			name:     "invalid JSON",
			filePath: "invalid.json",
			want:     nil,
			wantErr:  true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := ReadJSON(test.filePath)
			if test.wantErr == true {
				require.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
