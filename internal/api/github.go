package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/letsmakecakes/github-trending-cli/internal/config"
	"github.com/letsmakecakes/github-trending-cli/internal/models"
)

const (
	githubAPIURL = "https://api.github.com/search/repositories"
	userAgent    = "github-trending-cli/1.0"
)

// Client represents a GitHub API client
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new GitHub API Client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: githubAPIURL,
	}
}

// FetchTrendingRepos fetches trending repositories from GitHub
func (c *Client) FetchTrendingRepos(cfg *config.Config) ([]models.Repository, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	query := c.buildQuery(cfg)
	repos, err := c.searchRepositories(query, cfg.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search repositories: %w", err)
	}

	return repos, nil
}

// buildQuery constructs the GitHub search query
func (c *Client) buildQuery(cfg *config.Config) string {
	daysAgo := cfg.GetDurationInDays()
	date := time.Now().AddDate(0, 0, -daysAgo).Format("2006-01-02")
	return fmt.Sprintf("created:>%s", date)
}

// searchRepositories performs the actual API call
func (c *Client) searchRepositories(query string, limit int) ([]models.Repository, error) {
	params := url.Values{}
	params.Add("q", query)
	params.Add("sort", "stars")
	params.Add("order", "desc")
	params.Add("per_page", fmt.Sprintf("%d", limit))

	reqURL := fmt.Sprintf("%s?%s", c.baseURL, params.Encode())

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github/v3+json")
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var searchResp models.SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return searchResp.Items, nil
}
