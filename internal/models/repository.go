package models

// Repository represents a GitHub repository
type Repository struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	HTMLURL         string `json:"html_url"`
	Stars           int    `json:"stargazers_count"`
	Forks           int    `json:"forks_count"`
	Language        string `json:"language"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Owner           Owner  `json:"owner"`
	OpenIssuesCount int    `json:"open_issues_count"`
}

// Owner represents the repository owner
type Owner struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
}

// SearchResponse represents GitHub API search response
type SearchResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}
