package plentylog

import (
	"context"
	"time"

	"github.com/rs/xid"
)

type Transaction struct {
	*Log
	id   string
	logs []Record
}

func (pl *Log) NewTransaction() *Transaction {
	t := Transaction{
		id: xid.New().String(),
	}

	t.Log = pl

	return &t
}

func (t *Transaction) Debug(message string, metadata Metadata) {
	t.addLog(levelDebug, message, metadata)
}

func (t *Transaction) Info(message string, metadata Metadata) {
	t.addLog(levelInfo, message, metadata)
}

func (t *Transaction) Warning(message string, metadata Metadata) {
	t.addLog(levelWarning, message, metadata)
}

func (t *Transaction) Error(message string, metadata Metadata) {
	t.addLog(levelError, message, metadata)
}

func (t *Transaction) Commit(ctx context.Context) {
	for _, log := range t.logs {
		t.Log.provider.Write(ctx, log)
	}
}

func (t *Transaction) Rollback() {
	t.id = ""
	t.logs = nil
}

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
