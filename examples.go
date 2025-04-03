package main

import (
	"context"
	"plentylog/plentylog"
	"time"
)

func main() {
	// simple cli logger
	pl := plentylog.NewLog(nil)

	pl.Debug("debug", plentylog.Metadata{
		"test":  "test",
		"test2": "test2",
	})

	pl.Debug("test message", plentylog.Metadata{
		"123": "123",
	})

	// tranzactions
	log := plentylog.NewLog(plentylog.NewProviderFile(&plentylog.ProviderFileOptions{Format: plentylog.FormatJSON}))

	go func() {
		tr := log.NewTransaction()

		tr.Debug("Debug message", plentylog.Metadata{"key": "value"})
		tr.Info("Info message", plentylog.Metadata{"key2": "value2"})

		tr.Commit(context.Background())
	}()

	go func() {
		tr := log.NewTransaction()

		tr.Debug("Debug message", plentylog.Metadata{"key": "value"})
		tr.Info("Info message", plentylog.Metadata{"key2": "value2"})

		tr.Commit(context.Background())
	}()

	time.Sleep(2 * time.Second)
}
