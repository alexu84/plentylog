package plentylog

import (
	"context"
	"time"
)

// Provider interface defines the methods that a logging provider must implement.
// It is used to write log records to different destinations.
// The provider is responsible for handling the actual writing of log records.
type Provider interface {
	Write(context.Context, Record)
}

// Log represents the main logging structure.
type Log struct {
	provider Provider
	config   *config
}

// LogOptions represents the options for creating a new Log instance.
// It includes a configuration file and a custom provider.
type LogOptions struct {
	ConfigFile string
	Provider   Provider
}

// Metadata represents additional information that can be attached to log records.
type Metadata map[string]any

// level represents the log level of a log record.
type level string

const (
	levelDebug   level = "DEBUG"
	levelInfo    level = "INFO"
	levelWarning level = "WARNING"
	levelError   level = "ERROR"
)

// Record represents a log record.
type Record struct {
	Timestamp     time.Time `json:"timestamp"`
	Level         level     `json:"level"`
	TransactionID string    `json:"transaction_id"`
	Message       string    `json:"message"`
	Metadata      Metadata  `json:"metadata"`
}

// NewLog creates a new Log instance.
// It takes a LogOptions struct as an argument.
// If the Provider field is nil, it will load the configuration from the specified file.
// If the ConfigFile field is empty, it defaults to "config.yml".
func NewLog(opts *LogOptions) (*Log, error) {
	pl := Log{}

	if opts == nil {
		opts = &LogOptions{}
	}

	if opts.Provider != nil {
		pl.provider = opts.Provider

		return &pl, nil
	}

	if opts.ConfigFile == "" {
		opts.ConfigFile = "config.yml"
	}

	c, err := loadConfig(opts.ConfigFile)
	if err != nil {
		return nil, err
	}

	if c != nil {
		switch c.InternalProvider {
		case "file":
			pl.provider = NewProviderFile(&ProviderFileOptions{Format: c.FileFormat})
		case "cli":
			pl.provider = NewProviderCLI()
		default:
			pl.provider = NewProviderCLI()
		}

		return &pl, nil
	}

	pl.provider = NewProviderCLI()

	return &pl, nil
}

// Debug creates a debug log record.
func (pl *Log) Debug(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelDebug, message, metadata)
}

// DebugWithContext creates a debug log record with a context.
func (pl *Log) DebugWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelDebug, message, metadata)
}

// Info creates an info log record.
func (pl *Log) Info(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelInfo, message, metadata)
}

// InfoWithContext creates an info log record with a context.
func (pl *Log) InfoWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelInfo, message, metadata)
}

// Warning creates a warning log record.
func (pl *Log) Warning(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelWarning, message, metadata)
}

// WarningWithContext creates a warning log record with a context.
func (pl *Log) WarningWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelWarning, message, metadata)
}

// Error creates an error log record.
func (pl *Log) Error(message string, metadata Metadata) {
	pl.writeLog(context.Background(), levelError, message, metadata)
}

// ErrorWithContext creates an error log record with a context.
func (pl *Log) ErrorWithContext(ctx context.Context, message string, metadata Metadata) {
	pl.writeLog(ctx, levelError, message, metadata)
}

// writeLog writes a log record to the provider.
func (pl *Log) writeLog(ctx context.Context, level level, message string, metadata Metadata) {
	log := Record{
		Message:   message,
		Level:     level,
		Timestamp: time.Now(),
		Metadata:  metadata,
	}

	pl.provider.Write(ctx, log)
}
