package plentylog

import (
	"context"
	"fmt"
)

type ProviderCLI struct{}

func NewProviderCLI() *ProviderCLI {
	return &ProviderCLI{}
}

func (p *ProviderCLI) Write(_ context.Context, l log) error {
	fmt.Println(textSerialization(l))

	return nil
}
