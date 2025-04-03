package plentylog

import (
	"reflect"
	"testing"
)

func TestTransaction_addLog(t *testing.T) {
	tests := []struct {
		name     string
		level    level
		message  string
		metadata Metadata
	}{
		{
			name:     "add debug log",
			level:    levelDebug,
			message:  "debug message",
			metadata: Metadata{"key1": "value1"},
		},
		{
			name:     "add info log",
			level:    levelInfo,
			message:  "info message",
			metadata: Metadata{"key2": "value2"},
		},
		{
			name:     "add warning log",
			level:    levelWarning,
			message:  "warning message",
			metadata: Metadata{"key3": "value3"},
		},
		{
			name:     "add error log",
			level:    levelError,
			message:  "error message",
			metadata: Metadata{"key4": "value4"},
		},
		{
			name:     "add log with empty metadata",
			level:    levelDebug,
			message:  "message with no metadata",
			metadata: Metadata{},
		},
	}

	pl := NewPlentyLog(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction := pl.NewTransaction()
			transaction.addLog(tt.level, tt.message, tt.metadata)

			if len(transaction.logs) != 1 {
				t.Fatalf("Expected 1 log, got %d", len(transaction.logs))
			}

			log := transaction.logs[0]

			if log.TransactionID != transaction.id {
				t.Errorf("Transaction ID does not match. got: %s, want: %s", log.TransactionID, transaction.id)
			}

			if log.Message != tt.message {
				t.Errorf("Message does not match. got: %s, want: %s", log.Message, tt.message)
			}

			if log.Level != tt.level {
				t.Errorf("Level does not match. got: %s, want: %s", log.Level, tt.level)
			}

			if !reflect.DeepEqual(log.Metadata, tt.metadata) {
				t.Errorf("Metadata does not match. got: %v, want: %v", log.Metadata, tt.metadata)
			}

			if log.Timestamp.IsZero() {
				t.Errorf("Timestamp is zero")
			}
		})
	}
}
