package plentylog

import (
	"testing"
)

func TestDebug(t *testing.T) {
	pl := NewPlentyLog(nil)

	pl.Debug("debug", Metadata{
		"test":  "test",
		"test2": "test2",
	})

	pl.Debug("test message", Metadata{
		"123": "123",
	})
}
