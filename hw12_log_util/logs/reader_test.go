package logs_test

import (
	"os"
	"testing"

	"github.com/goodelias/otusgo-basic/hw12_log_util/logs"
	"github.com/stretchr/testify/assert"
)

func TestReadLogs(t *testing.T) {
	tests := []struct {
		name      string
		fileSetup func(t *testing.T) string // функция для создания тестового файла
		expected  []string
		expectErr bool
	}{
		{
			name: "Valid file with multiple lines",
			fileSetup: func(t *testing.T) string {
				t.Helper()
				file, err := os.CreateTemp("", "valid_logs.txt")
				if err != nil {
					t.Fatalf("не удалось создать временный файл: %v", err)
				}
				content := "line 1\nline 2\nline 3\n"
				_, _ = file.WriteString(content)
				file.Close()
				return file.Name()
			},
			expected:  []string{"line 1", "line 2", "line 3"},
			expectErr: false,
		},
		{
			name: "Empty file",
			fileSetup: func(t *testing.T) string {
				t.Helper()
				file, err := os.CreateTemp("", "empty_logs.txt")
				if err != nil {
					t.Fatalf("не удалось создать временный файл: %v", err)
				}
				file.Close()
				return file.Name()
			},
			expected:  []string{},
			expectErr: false,
		},
		{
			name: "File does not exist",
			fileSetup: func(_ *testing.T) string {
				return "nonexistent_file.txt"
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name: "File with invalid lines",
			fileSetup: func(t *testing.T) string {
				t.Helper()
				file, err := os.CreateTemp("", "invalid_logs.txt")
				if err != nil {
					t.Fatalf("не удалось создать временный файл: %v", err)
				}
				content := "line 1\nline 2\ninvalid line without newline"
				_, _ = file.WriteString(content)
				file.Close()
				return file.Name()
			},
			expected:  []string{"line 1", "line 2", "invalid line without newline"},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := tt.fileSetup(t)

			lines, err := logs.ReadLogs(filePath)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, lines)

			if _, err := os.Stat(filePath); err == nil {
				_ = os.Remove(filePath)
			}
		})
	}
}
