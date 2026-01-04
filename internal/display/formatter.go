package display

import (
	"fmt"
	"strings"

	"github.com/letsmakecakes/github-trending-cli/internal/models"
)

const (
	separator = "─"
	lineWidth = 100
)

// Formatter handles the display of repositories
type Formatter struct{}

// NewFormatter creates a new Formatter
func NewFormatter() *Formatter {
	return &Formatter{}
}

// Format formats repositories for display
func (f *Formatter) Format(repos []models.Repository) string {
	if len(repos) == 0 {
		return "No trending repositories found."
	}

	var sb strings.Builder

	sb.WriteString(f.header())
	sb.WriteString("\n\n")

	for i, repo := range repos {
		sb.WriteString(f.formatRepository(i+1, repo))
		if i < len(repos)-1 {
			sb.WriteString("\n")
			sb.WriteString(strings.Repeat(separator, lineWidth))
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// header returns the formatted header
func (f *Formatter) header() string {
	return fmt.Sprintf("🌟 Trending Repositories\\n%s", strings.Repeat("═", lineWidth))
}

// formatRepository formats a single repository
func (f *Formatter) formatRepository(rank int, repo models.Repository) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("#%d %s\n", rank, repo.FullName))
	sb.WriteString(fmt.Sprintf("    URL: %s\n", repo.HTMLURL))

	if repo.Description != "" {
		desc := repo.Description
		if len(desc) > 80 {
			desc = desc[:77] + "..."
		}
		sb.WriteString(fmt.Sprintf("    Description: %s\n", desc))
	}

	sb.WriteString(fmt.Sprintf("    ⭐ Stars: %s", f.formatNumber(repo.Stars)))

	if repo.Language != "" {
		sb.WriteString(fmt.Sprintf(" | 💻 Language: %s", repo.Language))
	}

	sb.WriteString(fmt.Sprintf(" | 🔱 Forks: %s", f.formatNumber(repo.Forks)))

	return sb.String()
}

// formatNumber formats a number with a thousand separators
func (f *Formatter) formatNumber(n int) string {
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}
	if n < 1000000 {
		return fmt.Sprintf("%.1fk", float64(n)/1000)
	}
	return fmt.Sprintf("%.1fM", float64(n)/1000000)
}
