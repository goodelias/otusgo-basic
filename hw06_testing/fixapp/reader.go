package fixapp

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJSON(filePath string) ([]Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s %w", filePath, err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s %w", filePath, err)
	}

	var data []Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s %w", filePath, err)
	}

	return data, nil
}
