package output

import (
	"fmt"
	"os"
	"strings"
)

// WriteStats выводит статистику либо в консоль, либо в файл.
func WriteStats(stats map[string]int, outputPath string) error {
	var result strings.Builder
	result.WriteString("--- Статистика по уровням логов ---\n")
	for level, count := range stats {
		result.WriteString(fmt.Sprintf("%s: %d\n", level, count))
	}

	if outputPath != "" {
		err := os.WriteFile(outputPath, []byte(result.String()), 0o600)
		if err != nil {
			return fmt.Errorf("не удалось записать в файл: %w", err)
		}
		fmt.Printf("Статистика записана в файл: %s\n", outputPath)
	} else {
		fmt.Println(result.String())
	}

	return nil
}
