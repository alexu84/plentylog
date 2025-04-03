package plentylog

import (
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	pl := NewPlentyLog(nil)

	err := pl.Debug(PlentyLogMetadata{
		"test":  "test",
		"test2": "test2",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = pl.Debug(PlentyLogMetadata{
		"123": "123",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	time.Sleep(1 * time.Second)
}
