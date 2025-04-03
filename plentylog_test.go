package plentylog

import (
	"testing"
)

func TestDebug(t *testing.T) {
	pl := NewPlentyLog(nil)

	err := pl.Debug("debug", Metadata{
		"test":  "test",
		"test2": "test2",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = pl.Debug("test message", Metadata{
		"123": "123",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
