package plentylog

import "fmt"

type ProviderCLI struct{}

func NewProviderCLI(logs chan plentyLog) *ProviderCLI {
	fmt.Printf("Starting CLI provider...\n")

	go func() {
		for log := range logs {
			fmt.Printf("%s [%s] %s: %v\n",
				log.timestamp.Format("2006-01-02 15:04:05"),
				log.transactionID,
				log.level,
				log.metadata,
			)
		}
	}()

	return &ProviderCLI{}
}

func (p *ProviderCLI) Write() error {
	return nil
}
