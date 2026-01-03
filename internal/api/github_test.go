package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/letsmakecakes/github-trending-cli/internal/config"
	"github.com/letsmakecakes/github-trending-cli/internal/models"
)

func TestClient_FetchTrendingRepos(t *testing.T) {
	mockRepos := models.SearchResponse{
		TotalCount: 1,
		Items: []models.Repository{
			{
				Name:        "test-repo",
				FullName:    "user/test-repo",
				Description: "A test repository",
				Stars:       100,
				Language:    "Go",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search/repositories" {
			t.Errorf("Expected path /search/repositories, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(mockRepos)
		if err != nil {
			t.Logf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient()
	client.baseURL = server.URL + "/search/repositories"

	cfg := &config.Config{
		Duration: config.DurationWeek,
		Limit:    10,
	}

	repos, err := client.FetchTrendingRepos(cfg)
	if err != nil {
		t.Fatalf("FetchTrendingRepos() error = %v", err)
	}

	if len(repos) != 1 {
		t.Errorf("Expected 1 repository, got %d", len(repos))
	}

	if repos[0].Name != "test-repo" {
		t.Errorf("Expected repo nae 'test-repo', got '%s'", repos[0].Name)
	}
}

func TestClient_FetchTrendingRepos_InvalidConfig(t *testing.T) {
	client := NewClient()
	cfg := &config.Config{
		Duration: "invalid",
		Limit:    10,
	}

	_, err := client.FetchTrendingRepos(cfg)
	if err == nil {
		t.Error("Expected error for invalid config, got nil")
	}
}

func TestClient_FetchTrendingRepos_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal Server Error"))
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient()
	client.baseURL = server.URL + "/search/repositories"

	cfg := &config.Config{
		Duration: config.DurationWeek,
		Limit:    10,
	}

	_, err := client.FetchTrendingRepos(cfg)
	if err == nil {
		t.Error("Expected error for API failure, got nil")
	}
}
