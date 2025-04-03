package plentylog

import (
	"time"
)

type plentyLogProvider interface {
	// Write(<-chan plentyLog) error
}

type PlentyLog struct {
	provider plentyLogProvider
	logs     chan plentyLog
}

type PlentyLogMetadata map[string]any

type plentyLogLevel string

const (
	plentyLogLevelDebug   plentyLogLevel = "DEBUG"
	plentyLogLevelInfo    plentyLogLevel = "INFO"
	plentyLogLevelWarning plentyLogLevel = "WARNING"
	plentyLogLevelError   plentyLogLevel = "ERROR"
)

type plentyLog struct {
	transactionID string
	level         plentyLogLevel
	timestamp     time.Time
	metadata      PlentyLogMetadata
}

func NewPlentyLog(provider *plentyLogProvider) *PlentyLog {
	pl := PlentyLog{
		logs: make(chan plentyLog),
	}

	if provider != nil {
		pl.provider = provider
	} else {
		pl.provider = NewProviderCLI(pl.logs)
	}

	return &pl
}

func (pl *PlentyLog) Debug(metadata PlentyLogMetadata) error {
	log := plentyLog{
		level:     plentyLogLevelDebug,
		timestamp: time.Now(),
		metadata:  metadata,
	}

	pl.logs <- log

	return nil
}
