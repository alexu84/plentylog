# plentylog

The plentylog package provides a flexible logging solution in Go, allowing logs to be written to different providers such as the command line interface or a file. It supports structured logging with metadata and transaction IDs.

## Key components:

- [plentylog/config.go](plentylog/config.go) - Defines the configuration structure for the logging system, including options for different providers and formats.
- [plentylog/plentylog.go](plentylog/plentylog.go) - Defines the core logging functionality, including the Log struct, Provider interface, different log levels, and methods for writing logs with and without context.
- [plentylog/provider_cli.go](plentylog/provider_cli.go) -  Implements a Provider that writes logs to the command line interface, using the textSerialization function to format the log message.
- [plentylog/provider_file.go](plentylog/provider_file.go) - Implements a Provider that writes logs to a file in either JSON or text format. It uses channels and goroutines for asynchronous writing and error handling.
- [plentylog/serialization.go](plentylog/serialization.go) - Contains functions for serializing log messages into different formats, including textSerialization and jsonSerialization.
- [plentylog/transaction.go](plentylog/transaction.go) - Introduces the concept of transactions, allowing multiple logs to be grouped and committed or rolled back together.
- [examples.go](examples.go) - Provides usage examples, demonstrating how to use the package with different providers and transactions.

## Example usage:

Simple logger with internal cli provider and config from file

```go
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
```

Transactions with internal file provider

```go
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
```

External elastic search provider

```go
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
```

## UML diagrams to illustrate the main classes and their relationships

![Class diagram](class_diagram.png)

![Sequence](sequence.png)
