package plentylog

import "testing"

func TestTransactionDebug(t *testing.T) {
	pl := NewPlentyLog(nil)

	tr := pl.NewTransaction()

	tr.Debug(PlentyLogMetadata{"test": "test"})
	tr.Debug(PlentyLogMetadata{"test2": "test2"})

	tr.Commit()

	tr = pl.NewTransaction()

	tr.Debug(PlentyLogMetadata{"123": "123"})
	tr.Debug(PlentyLogMetadata{"234": "234"})

	tr.Commit()
}
