package cli

import (
	"flag"
	"fmt"

	"github.com/letsmakecakes/github-trending-cli/internal/api"
	"github.com/letsmakecakes/github-trending-cli/internal/config"
	"github.com/letsmakecakes/github-trending-cli/internal/display"
)

// Execute runs the CLI application
func Execute() error {
	cfg, err := parseFlags()
	if err != nil {
		return err
	}

	client := api.NewClient()
	repos, err := client.FetchTrendingRepos(cfg)
	if err != nil {
		return fmt.Errorf("failed to fetch trending repositories: %w", err)
	}

	formatter := display.NewFormatter()
	output := formatter.Format(repos)
	fmt.Println(output)

	return nil
}

// parseFlags parses command-line flags
func parseFlags() (*config.Config, error) {
	duration := flag.String("duration", "week", "Time range: day, week, month, year")
	limit := flag.Int("limit", 10, "Number of repositories to display (1-100)")
	flag.Parse()

	cfg := &config.Config{
		Duration: config.Duration(*duration),
		Limit:    *limit,
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
