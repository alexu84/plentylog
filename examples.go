package main

import (
	"context"
	"plentylog/elasticprovider"
	"plentylog/plentylog"
	"time"
)

func main() {
	// simple cli logger
	pl, err := plentylog.NewLog(nil)
	if err != nil {
		panic(err)
	}

	pl.Debug("debug", plentylog.Metadata{
		"test":  "test",
		"test2": "test2",
	})

	pl.Debug("test message", plentylog.Metadata{
		"123": "123",
	})

	// transactions
	provider := plentylog.NewProviderFile(&plentylog.ProviderFileOptions{Format: plentylog.FormatJSON})

	log, err := plentylog.NewLog(&plentylog.LogOptions{Provider: provider})
	if err != nil {
		panic(err)
	}

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

	// elasticsearch external provider
	esProvider, err := elasticprovider.NewElasticProvider(nil)
	if err != nil {
		panic(err)
	}

	esLog, err := plentylog.NewLog(&plentylog.LogOptions{Provider: esProvider})
	if err != nil {
		panic(err)
	}

	esLog.Debug("Debug message", plentylog.Metadata{"key": "value"})
	esLog.Info("Info message", plentylog.Metadata{"key2": "value2"})
	esLog.Error("Error message", plentylog.Metadata{"key3": "value3"})
	esLog.Warning("Warn message", plentylog.Metadata{"key4": "value4"})

	// // view logs: https://localhost:9200/logs/_search
}
