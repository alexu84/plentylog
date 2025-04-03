package plentylog

import (
	"testing"
	"time"
)

func TestProviderCLI_serialize(t *testing.T) {
	testTime := time.Now()
	testTimeFormat := testTime.Format("2006-01-02 15:04:05")

	tests := []struct {
		name string
		log  log
		want string
	}{
		{
			name: "basic log",
			log: log{
				timestamp: testTime,
				level:     "INFO",
				message:   "test message",
				metadata:  Metadata{},
			},
			want: testTimeFormat + " INFO \"test message\" ",
		},
		{
			name: "log with metadata",
			log: log{
				timestamp: testTime,
				level:     "DEBUG",
				message:   "another message",
				metadata:  Metadata{"key1": "value1", "key2": 123},
			},
			want: testTimeFormat + " DEBUG \"another message\" key1: value1, key2: 123",
		},
		{
			name: "log with transaction ID",
			log: log{
				timestamp:     testTime,
				level:         "WARN",
				message:       "message with transaction id",
				transactionID: "123e4567-e89b-12d3-a456-426614174000",
				metadata:      Metadata{},
			},
			want: testTimeFormat + " WARN \"message with transaction id\" transaction id: 123e4567-e89b-12d3-a456-426614174000, ",
		},
		{
			name: "log with all fields",
			log: log{
				timestamp:     testTime,
				level:         "ERROR",
				message:       "full log",
				transactionID: "some-uuid",
				metadata:      Metadata{"a": 1, "b": "2", "c": true},
			},
			want: testTimeFormat + " ERROR \"full log\" transaction id: some-uuid, a: 1, b: 2, c: true",
		},
		{
			name: "empty message",
			log: log{
				timestamp: testTime,
				level:     "INFO",
				message:   "",
				metadata:  Metadata{},
			},
			want: testTimeFormat + " INFO \"\" ",
		},
		{
			name: "metadata with special characters",
			log: log{
				timestamp: testTime,
				level:     "INFO",
				message:   "test message",
				metadata:  Metadata{"key1": "value with , comma", "key2": "value with \" quote"},
			},
			want: testTimeFormat + " INFO \"test message\" key1: value with , comma, key2: value with \" quote",
		},
	}

	provider := NewProviderCLI()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.serialize(tt.log)
			if got != tt.want {
				t.Errorf("ProviderCLI.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
