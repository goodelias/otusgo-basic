package output

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteStats(t *testing.T) {
	tests := []struct {
		name        string
		stats       map[string]int
		outputPath  string
		expectError bool
	}{
		{
			name: "ValidStatsToFile",
			stats: map[string]int{
				"INFO":  10,
				"ERROR": 5,
			},
			outputPath:  "test_output.txt",
			expectError: false,
		},
		{
			name: "ValidStatsToConsole",
			stats: map[string]int{
				"DEBUG": 3,
				"WARN":  2,
			},
			outputPath:  "",
			expectError: false,
		},
		{
			name:        "EmptyStats",
			stats:       map[string]int{},
			outputPath:  "empty_output.txt",
			expectError: false,
		},
		{
			name: "InvalidFilePath",
			stats: map[string]int{
				"CRITICAL": 1,
			},
			outputPath:  "/invalid/path/output.txt",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.outputPath != "" {
				defer func() {
					_ = os.Remove(tt.outputPath) // Удаляем файл после теста, если он был создан
				}()
			}

			err := WriteStats(tt.stats, tt.outputPath)

			if tt.expectError {
				assert.Error(t, err, "ожидалась ошибка, но её не произошло")
			} else {
				assert.NoError(t, err, "ошибка не ожидалась, но произошла")

				if tt.outputPath != "" {
					// Проверяем, что файл был создан и содержит ожидаемые данные
					data, readErr := os.ReadFile(tt.outputPath)
					assert.NoError(t, readErr, "ошибка при чтении файла")
					assert.Contains(t, string(data), "--- Статистика по уровням логов ---", "неверное содержимое файла")
					for level, count := range tt.stats {
						assert.Contains(t, string(data), level, "не хватает уровня лога в файле")
						assert.Contains(t, string(data), fmt.Sprint(count), "не хватает значения уровня лога в файле")
					}
				}
			}
		})
	}
}
