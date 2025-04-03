package plentylog

import "fmt"

type ProviderCLI struct{}

func NewProviderCLI() *ProviderCLI {
	return &ProviderCLI{}
}

func (p *ProviderCLI) Write(log plentyLog) error {
	fmt.Printf("%s [%s] %s: %v\n",
		log.timestamp.Format("2006-01-02 15:04:05"),
		log.transactionID,
		log.level,
		log.metadata,
	)

	return nil
}
