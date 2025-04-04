package plentylog

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	f, err := os.CreateTemp("", "configuration_test_*.yml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	defer os.Remove(f.Name())

	tests := []struct {
		name       string
		fileData   string
		shouldFail bool
		expected   *config
	}{
		{
			name:       "valid config",
			fileData:   "internalProvider: cli\nfileFormat: json\n",
			shouldFail: false,
			expected: &config{
				InternalProvider: configProviderCLI,
				FileFormat:       FormatJSON,
			},
		},
		{
			name:       "valid config with file provider",
			fileData:   "internalProvider: file\nfileFormat: json\n",
			shouldFail: false,
			expected: &config{
				InternalProvider: configProviderFile,
				FileFormat:       FormatJSON,
			},
		},
		{
			name:       "invalid yaml",
			fileData:   "x",
			shouldFail: true,
			expected:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = f.Seek(0, 0)
			if err != nil {
				t.Fatalf("Failed to seek temp file: %v", err)
			}

			err = f.Truncate(0)
			if err != nil {
				t.Fatalf("Failed to truncate temp file: %v", err)
			}

			_, err = f.Write([]byte(tt.fileData))
			if err != nil {
				t.Fatalf("Failed to write to temp file: %v", err)
			}

			config, err := loadConfig(f.Name())
			if (err != nil) != tt.shouldFail {
				t.Fatalf("Expected error: %v, got: %v", tt.shouldFail, err)
			}

			if err == nil && config != nil {
				if config.InternalProvider != tt.expected.InternalProvider {
					t.Errorf("Expected InternalProvider: %v, got: %v", tt.expected.InternalProvider, config.InternalProvider)
				}

				if config.FileFormat != tt.expected.FileFormat {
					t.Errorf("Expected FileFormat: %v, got: %v", tt.expected.FileFormat, config.FileFormat)
				}
			}
		})
	}
}
