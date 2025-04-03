package plentylog

import (
	"context"
	"fmt"
)

type ProviderCLI struct{}

func NewProviderCLI() *ProviderCLI {
	return &ProviderCLI{}
}

func (p *ProviderCLI) Write(_ context.Context, l log) {
	fmt.Println(textSerialization(l))
}
