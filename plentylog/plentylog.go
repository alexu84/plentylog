package plentylog

import (
	"context"
	"time"
)

type Provider interface {
	Write(context.Context, log)
}

type Log struct {
	provider Provider
}

type Metadata map[string]any

type level string

const (
	levelDebug   level = "DEBUG"
	levelInfo    level = "INFO"
	levelWarning level = "WARNING"
	levelError   level = "ERROR"
)

type log struct {
	Timestamp     time.Time `json:"timestamp"`
	Level         level     `json:"level"`
	TransactionID string    `json:"transaction_id"`
	Message       string    `json:"message"`
	Metadata      Metadata  `json:"metadata"`
}

func NewLog(provider Provider) *Log {
	pl := Log{}

	if provider != nil {
		pl.provider = provider
	} else {
		pl.provider = NewProviderCLI()
	}

	return &pl
}

func (pl *Log) Debug(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelDebug, message, metadata)
}

func (pl *Log) DebugWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelDebug, message, metadata)
}

func (pl *Log) Info(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelInfo, message, metadata)
}

func (pl *Log) InfoWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelInfo, message, metadata)
}

func (pl *Log) Warning(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelWarning, message, metadata)
}

func (pl *Log) WarningWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelWarning, message, metadata)
}

func (pl *Log) Error(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelError, message, metadata)
}

func (pl *Log) ErrorWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelError, message, metadata)
}

func (pl *Log) writeLog(ctx context.Context, level level, message string, metadata Metadata) {
	log := log{
		Message:   message,
		Level:     level,
		Timestamp: time.Now(),
		Metadata:  metadata,
	}

	pl.provider.Write(ctx, log)
}
