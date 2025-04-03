package plentylog

import (
	"context"
	"time"
)

type Provider interface {
	Write(context.Context, log) error
}

type PlentyLog struct {
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
	transactionID string
	message       string
	level         level
	timestamp     time.Time
	metadata      Metadata
}

func NewPlentyLog(provider Provider) *PlentyLog {
	pl := PlentyLog{}

	if provider != nil {
		pl.provider = provider
	} else {
		pl.provider = NewProviderCLI()
	}

	return &pl
}

func (pl *PlentyLog) Debug(message string, metadata Metadata) error {
	return pl.writeLog(context.Background(), levelDebug, message, metadata)
}

func (pl *PlentyLog) DebugWithContext(ctx context.Context, message string, metadata Metadata) error {
	return pl.writeLog(ctx, levelDebug, message, metadata)
}

func (pl *PlentyLog) Info(message string, metadata Metadata) error {
	return pl.writeLog(context.Background(), levelInfo, message, metadata)
}

func (pl *PlentyLog) InfoWithContext(ctx context.Context, message string, metadata Metadata) error {
	return pl.writeLog(ctx, levelInfo, message, metadata)
}

func (pl *PlentyLog) Warning(message string, metadata Metadata) error {
	return pl.writeLog(context.Background(), levelWarning, message, metadata)
}

func (pl *PlentyLog) WarningWithContext(ctx context.Context, message string, metadata Metadata) error {
	return pl.writeLog(ctx, levelWarning, message, metadata)
}

func (pl *PlentyLog) Error(message string, metadata Metadata) error {
	return pl.writeLog(context.Background(), levelError, message, metadata)
}

func (pl *PlentyLog) ErrorWithContext(ctx context.Context, message string, metadata Metadata) error {
	return pl.writeLog(ctx, levelError, message, metadata)
}

func (pl *PlentyLog) writeLog(ctx context.Context, level level, message string, metadata Metadata) error {
	log := log{
		message:   message,
		level:     level,
		timestamp: time.Now(),
		metadata:  metadata,
	}

	return pl.provider.Write(ctx, log)
}
