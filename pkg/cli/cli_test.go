package cli

import (
	"flag"
	"os"
	"testing"

	"github.com/letsmakecakes/github-trending-cli/internal/config"
)

func TestParseFlags(t *testing.T) {
	// Save original args and restore after a test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		duration config.Duration
		limit    int
	}{
		{
			name:     "default values",
			args:     []string{"cmd"},
			wantErr:  false,
			duration: config.DurationWeek,
			limit:    10,
		},
		{
			name:     "custom values",
			args:     []string{"cmd", "-duration", "month", "-limit", "20"},
			wantErr:  false,
			duration: config.DurationMonth,
			limit:    20,
		},
		{
			name:    "invalid duration",
			args:    []string{"cmd", "-duration", "invalid"},
			wantErr: true,
		},
		{
			name:    "invalid limit",
			args:    []string{"cmd", "-limit", "0"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flag package for each test
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			os.Args = tt.args
			cfg, err := parseFlags()

			if (err != nil) != tt.wantErr {
				t.Errorf("parseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if cfg.Duration != tt.duration {
					t.Errorf("Duration = %v, want %v", cfg.Duration, tt.duration)
				}
				if cfg.Limit != tt.limit {
					t.Errorf("Limit = %v, want %v", cfg.Limit, tt.limit)
				}
			}
		})
	}
}
