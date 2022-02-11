package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	client *http.Client
}

func NewHandler() *Handler {
	return &Handler{
		client: &http.Client{},
	}
}

// Github API
const githubAPI = "https://api.github.com"

type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	URL               string `json:"url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	HTMLURL           string `json:"html_url"`
	ReposURL          string `json:"repos_url"`
}
type Repository struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Owner       Owner     `json:"owner"`
	Private     bool      `json:"private"`
	HTMLURL     string    `json:"html_url"`
	Description string    `json:"description"`
	Fork        bool      `json:"fork"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PushedAt    time.Time `json:"pushed_at"`
	Homepage    string    `json:"homepage"`
	Size        int       `json:"size"`
	Language    string    `json:"language"`
	Visibility  string    `json:"visibility"`
}

type SearchResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

func (h *Handler) GetRepositories(search, sort, ignore string, page int) ([]Repository, error) {
	query := fmt.Sprintf("q=%s+in:name&page=%d&sort=name&order=%s", search, page, sort)
	// request setup
	url := fmt.Sprintf("%s/search/repositories?%s", githubAPI, query)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request, err: %s", err.Error())
	}
	// headers setup
	req.Header.Add("accept", "application/vnd.github.v3+json")
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do request, err: %s", err.Error())
	}
	searchResponse := &SearchResponse{}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting data, status code:%d", resp.StatusCode)
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&searchResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding body, err: %s", err.Error())
	}
	if ignore != "" {
		items := make([]Repository, 0)
    for _, repo := range searchResponse.Items {
      if !strings.Contains(repo.Name, ignore){
        items = append(items, repo)
      }
    }
    return items, nil
	}
	return searchResponse.Items, nil
}
