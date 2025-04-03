package plentylog

import (
	"time"

	"github.com/rs/xid"
)

type PlentyLogTransaction struct {
	*PlentyLog
	id   string
	logs []plentyLog
}

func (pl *PlentyLog) NewTransaction() *PlentyLogTransaction {
	plt := PlentyLogTransaction{
		id: xid.New().String(),
	}

	plt.PlentyLog = pl

	return &plt
}

func (plt *PlentyLogTransaction) Debug(metadata PlentyLogMetadata) error {
	log := plentyLog{
		transactionID: plt.id,
		level:         plentyLogLevelDebug,
		timestamp:     time.Now(),
		metadata:      metadata,
	}

	plt.logs = append(plt.logs, log)

	return nil
}

func (plt *PlentyLogTransaction) Commit() error {
	for _, log := range plt.logs {
		plt.PlentyLog.provider.Write(log)
	}

	return nil
}

func (plt *PlentyLogTransaction) Rollback() error {
	plt.logs = nil

	return nil
}
