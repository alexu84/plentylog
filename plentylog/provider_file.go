package plentylog

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

// ProviderFile is a provider for writing logs to a file.
// It implements the Provider interface from the plentylog package
type ProviderFile struct {
	opts   ProviderFileOptions
	mu     sync.Mutex
	logs   chan Record
	errors chan error
}

// ProviderFileOptions are the options for the ProviderFile
// It contains the file path and the format of the logs.
type ProviderFileOptions struct {
	FilePath string
	Format   format
}

type format string

const (
	FormatJSON format = "json"
	FormatText format = "text"
)

// NewProviderFile creates a new ProviderFile instance.
// It takes an optional ProviderFileOptions struct.
// If the FilePath field is empty, it defaults to "log.txt".
// If the Format field is empty, it defaults to text format.
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
		logs:   make(chan Record, 100),
		errors: make(chan error, 100),
	}

	// writeLogs in a separate goroutine
	// to avoid blocking the main thread
	// and to allow for concurrent writes
	go pf.writeLogs()

	// displayErrors in a separate goroutine
	go pf.displayErrors()

	return &pf
}

// Write writes a log record to a channel.
func (p *ProviderFile) Write(_ context.Context, l Record) {
	p.logs <- l
}

// writeLogs writes logs to a file.
// It listens for log records on the logs channel
// and writes them to the specified file.
// It uses a ticker to periodically check for new logs
// and write them to the file.
// It also handles errors by sending them to the errors channel.
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
