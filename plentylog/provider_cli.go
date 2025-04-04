package plentylog

import (
	"context"
	"fmt"
)

// ProviderCLI is a provider for writing logs to the CLI.
type ProviderCLI struct{}

// NewProviderCLI creates a new ProviderCLI instance.
func NewProviderCLI() *ProviderCLI {
	return &ProviderCLI{}
}

// Write writes a log record to the CLI.
func (p *ProviderCLI) Write(_ context.Context, r Record) {
	fmt.Println(textSerialization(r))
}
