package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config хранит параметры для анализа логов.
type Config struct {
	FilePath string
	LogLevel string
	Output   string
}

// LoadConfig загружает конфигурацию из флагов и переменных окружения.
func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки .env файла")
	}

	filePath := flag.String("file", "", "logs file path")
	logLevel := flag.String("level", "info", "logs level for analyzing")
	output := flag.String("output", "", "statistic file path")
	flag.Parse()

	if *filePath == "" {
		if envFilePath := os.Getenv("LOG_ANALYZER_FILE"); envFilePath != "" {
			filePath = &envFilePath
		}
	}

	if *logLevel == "info" {
		if envLogLevel := os.Getenv("LOG_ANALYZER_LEVEL"); envLogLevel != "" {
			logLevel = &envLogLevel
		}
	}

	if *output == "" {
		if envOutput := os.Getenv("LOG_ANALYZER_OUTPUT"); envOutput != "" {
			output = &envOutput
		}
	}

	if *filePath == "" {
		return Config{}, errors.New("не указан путь к файлу логов")
	}

	return Config{
		FilePath: *filePath,
		LogLevel: *logLevel,
		Output:   *output,
	}, nil
}
