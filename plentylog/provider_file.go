package plentylog

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type ProviderFile struct {
	opts   ProviderFileOptions
	mu     sync.Mutex
	logs   chan log
	errors chan error
}

type ProviderFileOptions struct {
	FilePath string
	Format   format
}

type format string

const (
	FormatJSON format = "json"
	FormatText format = "text"
)

func NewProviderFile(opts *ProviderFileOptions) *ProviderFile {
	if opts == nil {
		opts = &ProviderFileOptions{}
	}

	if opts.Format == "" {
		opts.Format = FormatText
	}

	if opts.FilePath == "" {
		switch opts.Format {
		case FormatJSON:
			opts.FilePath = "log.json"
		case FormatText:
			opts.FilePath = "log.txt"
		}
	}

	pf := ProviderFile{
		opts:   *opts,
		mu:     sync.Mutex{},
		logs:   make(chan log, 100),
		errors: make(chan error, 100),
	}

	go pf.writeLogs()
	go pf.displayErrors()

	return &pf
}

func (p *ProviderFile) Write(_ context.Context, l log) {
	p.logs <- l
}

func (p *ProviderFile) writeLogs() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case l := <-p.logs:
			p.mu.Lock()
			file, err := os.OpenFile(p.opts.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				p.errors <- err

				p.mu.Unlock()

				continue
			}

			var sd string
			switch p.opts.Format {
			case FormatJSON:
				s, err := jsonSerialization(l)
				if err != nil {
					p.errors <- err

					file.Close()
					p.mu.Unlock()

					continue
				}

				sd = *s
			case FormatText:
				sd = textSerialization(l)
			default:
				p.errors <- errors.New("unsupported serialization format")

				file.Close()
				p.mu.Unlock()

				continue
			}

			_, err = file.WriteString(sd + "\n")
			if err != nil {
				p.errors <- err
			}

			file.Close()
			p.mu.Unlock()
		case <-ticker.C:
			continue
		}
	}
}

func (p *ProviderFile) displayErrors() {
	for err := range p.errors {
		fmt.Println("Provider file ERROR: " + err.Error())
	}
}
