package logs

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLogs читает лог-файл и возвращает список строк.
func ReadLogs(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("не удалось закрыть файл: %v\n", err)
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	if lines == nil {
		return []string{}, nil
	}

	return lines, nil
}
