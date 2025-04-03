package plentylog

import "testing"

func TestTransactionDebug(t *testing.T) {
	pl := NewPlentyLog(nil)

	tr := pl.NewTransaction()

	tr.Debug("some error", Metadata{"test": "test"})
	tr.Debug("some error 2", Metadata{"test2": "test2"})

	tr.Commit(t.Context())

	tr = pl.NewTransaction()

	tr.Debug("more error", Metadata{"123": "123"})
	tr.Debug("more error 2", Metadata{"234": "234"})

	tr.Commit(t.Context())
}
