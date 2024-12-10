package logs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goodelias/otusgo-basic/hw12_log_util/models"
)

// AnalyzeLogs анализирует логи и возвращает статистику по уровням.
func AnalyzeLogs(logs []string, logLevel string) map[string]int {
	logLevelsCount := make(map[string]int)

	for _, line := range logs {
		var logEntry models.LogEntry
		if err := json.Unmarshal([]byte(line), &logEntry); err != nil {
			fmt.Printf("Ошибка при разборе строки лога: %v, строка: %s\n", err, line)
			continue
		}

		if logLevel == "" || strings.EqualFold(logEntry.Level, logLevel) {
			logLevelsCount[logEntry.Level]++
		}
	}

	return logLevelsCount
}
