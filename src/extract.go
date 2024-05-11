package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Repository struct {
	Name        string `json:"name"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	Fork        bool   `json:"fork"`
	Archived    bool   `json:"archived"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Description string `json:"description"`
}

func getPublicRepos(username string) ([]Repository, error) {
	allRepos := make([]Repository, 0)
	page := 1
	for {
		url := fmt.Sprintf("https://api.github.com/users/%s/repos?page=%d&per_page=100", username, page)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("Failed to fetch data: %s", resp.Status)
		}
		var repos []Repository
		if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()
		if len(repos) == 0 {
			break
		}
		allRepos = append(allRepos, repos...)
		page++
		time.Sleep(time.Second)
	}
	return allRepos, nil
}

func main() {
	username := "vilaca"
	repos, err := getPublicRepos(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, repo := range repos {
		fmt.Printf("%d %d %t %t %s %s %s %s\n", repo.Stars, repo.Forks, repo.Fork, repo.Archived, repo.Name, repo.Created, repo.Updated, repo.Description)
	}
}
