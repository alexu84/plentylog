package plentylog

import (
	"context"
	"testing"
	"time"
)

func TestProviderTextFile(t *testing.T) {
	log := NewPlentyLog(NewProviderFile(&ProviderFileOptions{Format: formatJSON}))

	go func() {
		tr := log.NewTransaction()

		tr.Debug("Debug message", Metadata{"key": "value"})
		tr.Info("Info message", Metadata{"key2": "value2"})

		err := tr.Commit(context.Background())
		if err != nil {
			t.Errorf("Error committing transaction: %v", err)
		}
	}()

	go func() {
		tr := log.NewTransaction()

		tr.Debug("Debug message", Metadata{"key": "value"})
		tr.Info("Info message", Metadata{"key2": "value2"})

		err := tr.Commit(context.Background())
		if err != nil {
			t.Errorf("Error committing transaction: %v", err)
		}
	}()

	time.Sleep(2 * time.Second)
}
