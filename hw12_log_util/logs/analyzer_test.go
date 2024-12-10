package logs_test

import (
	"testing"

	"github.com/goodelias/otusgo-basic/hw12_log_util/logs"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzeLogs(t *testing.T) {
	tests := []struct {
		name     string
		logs     []string
		logLevel string
		expected map[string]int
	}{
		{
			name: "All levels counted",
			logs: []string{
				`{"timestamp":"2024-01-01T00:00:00Z","level":"info","message":"Test message 1"}`,
				`{"timestamp":"2024-01-01T01:00:00Z","level":"error","message":"Test message 2"}`,
				`{"timestamp":"2024-01-01T02:00:00Z","level":"debug","message":"Test message 3"}`,
			},
			logLevel: "",
			expected: map[string]int{
				"info":  1,
				"error": 1,
				"debug": 1,
			},
		},
		{
			name: "Filter by specific level",
			logs: []string{
				`{"timestamp":"2024-01-01T00:00:00Z","level":"info","message":"Test message 1"}`,
				`{"timestamp":"2024-01-01T01:00:00Z","level":"info","message":"Test message 2"}`,
				`{"timestamp":"2024-01-01T02:00:00Z","level":"error","message":"Test message 3"}`,
			},
			logLevel: "info",
			expected: map[string]int{
				"info": 2,
			},
		},
		{
			name: "Ignore invalid logs",
			logs: []string{
				`{"timestamp":"2024-01-01T00:00:00Z","level":"info","message":"Test message 1"}`,
				`Invalid log line`,
				`{"timestamp":"2024-01-01T02:00:00Z","level":"error","message":"Test message 2"}`,
			},
			logLevel: "",
			expected: map[string]int{
				"info":  1,
				"error": 1,
			},
		},
		{
			name:     "Empty logs",
			logs:     []string{},
			logLevel: "",
			expected: map[string]int{},
		},
		{
			name: "Case insensitive level filter",
			logs: []string{
				`{"timestamp":"2024-01-01T00:00:00Z","level":"INFO","message":"Test message 1"}`,
				`{"timestamp":"2024-01-01T01:00:00Z","level":"info","message":"Test message 2"}`,
				`{"timestamp":"2024-01-01T02:00:00Z","level":"Error","message":"Test message 3"}`,
			},
			logLevel: "info",
			expected: map[string]int{
				"INFO": 1,
				"info": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := logs.AnalyzeLogs(tt.logs, tt.logLevel)
			assert.Equal(t, tt.expected, result)
		})
	}
}
