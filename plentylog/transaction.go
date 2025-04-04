package plentylog

import (
	"context"
	"time"

	"github.com/rs/xid"
)

// Transaction represents a logging transaction.
type Transaction struct {
	*Log
	id   string
	logs []Record
}

// NewTransaction creates a new transaction for logging.
func (pl *Log) NewTransaction() *Transaction {
	t := Transaction{
		id: xid.New().String(),
	}

	t.Log = pl

	return &t
}

// Debug creates a debug log record.
func (t *Transaction) Debug(message string, metadata Metadata) {
	t.addLog(levelDebug, message, metadata)
}

// Info creates an info log record.
func (t *Transaction) Info(message string, metadata Metadata) {
	t.addLog(levelInfo, message, metadata)
}

// Warning creates a warning log record.
func (t *Transaction) Warning(message string, metadata Metadata) {
	t.addLog(levelWarning, message, metadata)
}

// Error creates an error log record.
func (t *Transaction) Error(message string, metadata Metadata) {
	t.addLog(levelError, message, metadata)
}

// Commit writes all logs in the transaction to the provider.
func (t *Transaction) Commit(ctx context.Context) {
	for _, log := range t.logs {
		t.Log.provider.Write(ctx, log)
	}
}

// Rollback discards all logs in the transaction.
// It resets the transaction ID and clears the logs.
func (t *Transaction) Rollback() {
	t.id = ""
	t.logs = nil
}

// addLog adds a log record to the transaction.
// It creates a new log record with the given level, message, and metadata,
func (t *Transaction) addLog(lv level, message string, metadata Metadata) {
	l := Record{
		TransactionID: t.id,
		Message:       message,
		Level:         lv,
		Timestamp:     time.Now(),
		Metadata:      metadata,
	}

	t.logs = append(t.logs, l)
}
