package plentylog

import (
	"testing"
	"time"
)

func TestTextSerialization(t *testing.T) {
	testTime := time.Now()
	testTimeFormat := testTime.Format("2006-01-02 15:04:05")

	tests := []struct {
		name string
		log  Record
		want string
	}{
		{
			name: "basic log",
			log: Record{
				Timestamp: testTime,
				Level:     "INFO",
				Message:   "test message",
				Metadata:  Metadata{},
			},
			want: testTimeFormat + " INFO \"test message\" ",
		},
		{
			name: "log with metadata",
			log: Record{
				Timestamp: testTime,
				Level:     "DEBUG",
				Message:   "another message",
				Metadata:  Metadata{"key1": "value1", "key2": 123},
			},
			want: testTimeFormat + " DEBUG \"another message\" key1: value1, key2: 123",
		},
		{
			name: "log with transaction ID",
			log: Record{
				Timestamp:     testTime,
				Level:         "WARN",
				Message:       "message with transaction id",
				TransactionID: "123e4567-e89b-12d3-a456-426614174000",
				Metadata:      Metadata{},
			},
			want: testTimeFormat + " WARN \"message with transaction id\" transaction id: 123e4567-e89b-12d3-a456-426614174000, ",
		},
		{
			name: "log with all fields",
			log: Record{
				Timestamp:     testTime,
				Level:         "ERROR",
				Message:       "full log",
				TransactionID: "some-uuid",
				Metadata:      Metadata{"a": 1, "b": "2", "c": true},
			},
			want: testTimeFormat + " ERROR \"full log\" transaction id: some-uuid, a: 1, b: 2, c: true",
		},
		{
			name: "empty message",
			log: Record{
				Timestamp: testTime,
				Level:     "INFO",
				Message:   "",
				Metadata:  Metadata{},
			},
			want: testTimeFormat + " INFO \"\" ",
		},
		{
			name: "metadata with special characters",
			log: Record{
				Timestamp: testTime,
				Level:     "INFO",
				Message:   "test message",
				Metadata:  Metadata{"key1": "value with , comma", "key2": "value with \" quote"},
			},
			want: testTimeFormat + " INFO \"test message\" key1: value with , comma, key2: value with \" quote",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := textSerialization(tt.log)
			if got != tt.want {
				t.Errorf("ProviderCLI.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
