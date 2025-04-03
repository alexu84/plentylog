package plentylog

import (
	"context"
	"fmt"
)

type ProviderCLI struct{}

func NewProviderCLI() *ProviderCLI {
	return &ProviderCLI{}
}

func (p *ProviderCLI) Write(_ context.Context, r Record) {
	fmt.Println(textSerialization(r))
}
