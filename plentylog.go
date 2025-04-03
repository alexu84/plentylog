package plentylog

import (
	"time"
)

type plentyLogProvider interface {
	Write(plentyLog) error
}

type PlentyLog struct {
	provider plentyLogProvider
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

func NewPlentyLog(provider plentyLogProvider) *PlentyLog {
	pl := PlentyLog{}

	if provider != nil {
		pl.provider = provider
	} else {
		pl.provider = NewProviderCLI()
	}

	return &pl
}

func (pl *PlentyLog) Debug(metadata PlentyLogMetadata) error {
	log := plentyLog{
		level:     plentyLogLevelDebug,
		timestamp: time.Now(),
		metadata:  metadata,
	}

	return pl.provider.Write(log)
}
