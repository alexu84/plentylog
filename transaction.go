package plentylog

import (
	"github.com/rs/xid"
)

type PlentyLogTransaction struct {
	// plentyLog
	id string
}

func (pl *PlentyLog) NewTransaction() *PlentyLogTransaction {
	return &PlentyLogTransaction{
		id: xid.New().String(),
	}
}

func (plt *PlentyLogTransaction) Debug(metadata PlentyLogMetadata) error {
	// log := plentyLog{
	// 	transactionID: plt.id,
	// 	level:         plentyLogLevelDebug,
	// 	timestamp:     time.Now(),
	// 	metadata:      []PlentyLogMetadata{metadata},
	// }

	// plt.plentyLog.metadata = append(plt.plentyLog.metadata, log)

	return nil
}

func (plt *PlentyLogTransaction) Commit() error {
	return nil
}

func (plt *PlentyLogTransaction) Rollback() error {

	return nil
}
