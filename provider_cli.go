package plentylog

type ProviderCLI struct{}

func NewProviderCLI(logs <-chan plentyLog) *ProviderCLI {
	return &ProviderCLI{}
}

func (p *ProviderCLI) Write() error {
	return nil
}
