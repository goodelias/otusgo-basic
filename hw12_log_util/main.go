package main

import (
	"fmt"

	"github.com/goodelias/otusgo-basic/hw12_log_util/config"
	"github.com/goodelias/otusgo-basic/hw12_log_util/logs"
	"github.com/goodelias/otusgo-basic/hw12_log_util/output"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v\n", err)
		return
	}

	logsData, err := logs.ReadLogs(cfg.FilePath)
	if err != nil {
		fmt.Printf("Ошибка чтения логов: %v\n", err)
		return
	}

	stats := logs.AnalyzeLogs(logsData, cfg.LogLevel)

	err = output.WriteStats(stats, cfg.Output)
	if err != nil {
		fmt.Printf("Ошибка вывода статистики: %v\n", err)
	}
}
