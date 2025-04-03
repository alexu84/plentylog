package plentylog

import (
	"bytes"
	"context"
	"fmt"
)

type ProviderCLI struct{}

func NewProviderCLI() *ProviderCLI {
	return &ProviderCLI{}
}

func (p *ProviderCLI) Write(_ context.Context, l log) error {
	fmt.Println(p.serialize(l))

	return nil
}

func (p *ProviderCLI) serialize(l log) string {
	var buffer bytes.Buffer

	buffer.WriteString(l.timestamp.Format("2006-01-02 15:04:05"))
	buffer.WriteString(" ")
	buffer.WriteString(string(l.level))
	buffer.WriteString(" \"")
	buffer.WriteString(l.message)
	buffer.WriteString("\" ")
	if l.transactionID != "" {
		buffer.WriteString("transaction id: ")
		buffer.WriteString(l.transactionID)
		buffer.WriteString(", ")
	}

	index := 0
	for key, value := range l.metadata {
		buffer.WriteString(key)
		buffer.WriteString(": ")
		buffer.WriteString(fmt.Sprintf("%v", value))
		if index < len(l.metadata)-1 {
			buffer.WriteString(", ")
		}
		index++
	}

	return buffer.String()
}
