package display

import (
	"strings"
	"testing"

	"github.com/letsmakecakes/github-trending-cli/internal/models"
)

func TestFormatter_Format(t *testing.T) {
	formatter := NewFormatter()

	tests := []struct {
		name     string
		repos    []models.Repository
		expected []string
	}{
		{
			name:     "empty repositories",
			repos:    []models.Repository{},
			expected: []string{"No trending repositories found."},
		},
		{
			name: "single repository",
			repos: []models.Repository{
				{
					Name:        "test-repo",
					FullName:    "user/test-repo",
					Description: "Test description",
					HTMLURL:     "https://github.com/user/test-repo",
					Stars:       1500,
					Language:    "Go",
					Forks:       50,
				},
			},
			expected: []string{"#1 user/test-repo", "1.5k", "Go"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatter.Format(tt.repos)
			for _, exp := range tt.expected {
				if !strings.Contains(result, exp) {
					t.Errorf("Expected output to contain '%s', got:\n%s", exp, result)
				}
			}
		})
	}
}

func TestFormatter_formatNumber(t *testing.T) {
	formatter := NewFormatter()

	tests := []struct {
		input    int
		expected string
	}{
		{500, "500"},
		{1500, "1.5k"},
		{1500000, "1.5M"},
		{999, "999"},
		{1000, "1.0k"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := formatter.formatNumber(tt.input)
			if result != tt.expected {
				t.Errorf("formatNumber(%d) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}
