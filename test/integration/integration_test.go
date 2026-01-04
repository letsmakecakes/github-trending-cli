package integration

import (
	"testing"
	"time"

	"github.com/letsmakecakes/github-trending-cli/internal/api"
	"github.com/letsmakecakes/github-trending-cli/internal/config"
)

func TestIntegration_FetchTrendingRepos(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := api.NewClient()

	tests := []struct {
		name     string
		config   *config.Config
		wantErr  bool
		minRepos int
	}{
		{
			name: "fetch weekly trending",
			config: &config.Config{
				Duration: config.DurationWeek,
				Limit:    5,
			},
			wantErr:  false,
			minRepos: 1,
		},
		{
			name: "fetch daily trending",
			config: &config.Config{
				Duration: config.DurationDay,
				Limit:    10,
			},
			wantErr:  false,
			minRepos: 1,
		},
		{
			name: "fetch monthly trending",
			config: &config.Config{
				Duration: config.DurationMonth,
				Limit:    15,
			},
			wantErr:  false,
			minRepos: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repos, err := client.FetchTrendingRepos(tt.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchTrendingRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(repos) < tt.minRepos {
					t.Errorf("Expected at least %d repos, got %d", tt.minRepos, len(repos))
				}

				// Verify the repository structure
				for i, repo := range repos {
					if repo.FullName == "" {
						t.Errorf("Repo %d has empty FullName", i)
					}
					if repo.HTMLURL == "" {
						t.Errorf("Repo %d has empty HTMLURL", i)
					}
					if repo.Stars < 0 {
						t.Errorf("Repo %d has negative stars", i)
					}
				}

				// Verify sorting by stars (descending)
				for i := 1; i < len(repos); i++ {
					if repos[i-1].Stars < repos[i].Stars {
						t.Errorf("Repos not sorted by stars: repo %d has %d stars, repo %d has %d stars",
							i-1, repos[i-1].Stars, i, repos[i].Stars)
					}
				}
			}
		})

		// Rate limiting courtesy
		time.Sleep(1 * time.Second)
	}
}

func TestIntegration_RateLimiting(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := api.NewClient()
	cfg := &config.Config{
		Duration: config.DurationWeek,
		Limit:    5,
	}

	// Make multiple requests to test rate limiting handling
	for i := 0; i < 3; i++ {
		_, err := client.FetchTrendingRepos(cfg)
		if err != nil {
			t.Logf("Request %d: %v", i+1, err)
		}
		time.Sleep(2 * time.Second)
	}
}
