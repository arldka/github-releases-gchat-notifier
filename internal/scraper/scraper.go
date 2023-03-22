package scraper

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"

	"github.com/arldka/github-releases-gchat-notifier/internal/models"
)

type Client struct {
	*github.Client
}

func NewClient() Client {

	token := os.Getenv("GH_ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	return Client{client}
}

func (client Client) GetReleases(repositories []models.Repository) ([]models.Repository, []models.Release) {

	var updatedRepositories = []models.Repository{}
	var releases = []models.Release{}

	for i, r := range repositories {
		// Split repository name into owner and name
		parts := strings.Split(r.Name, "/")
		if len(parts) != 2 {
			log.Fatalf("Invalid repository name: %s", r.Name)
		}
		owner, name := parts[0], parts[1]

		// Retrieve the latest release information for the repository
		release, _, err := client.Repositories.GetLatestRelease(context.Background(), owner, name)
		if err != nil {
			log.Fatalf("Failed to retrieve release information for %s: %s", r.Name, err)
		}

		if repositories[i].Tag != release.GetTagName() {
			repositories[i].Tag = release.GetTagName()
			repositories[i].Notified = false
			releases = append(releases, models.Release{
				Name:       release.GetName(),
				Tag:        release.GetTagName(),
				ReleaseURL: release.GetHTMLURL(),
				RepoName:   repositories[i].Name,
			})
			updatedRepositories = append(updatedRepositories, repositories[i])
		} else if !repositories[i].Notified {
			releases = append(releases, models.Release{Name: release.GetName(),
				Tag:        release.GetTagName(),
				ReleaseURL: release.GetHTMLURL(),
				RepoName:   repositories[i].Name,
			})
			updatedRepositories = append(updatedRepositories, repositories[i])
		}

	}

	return updatedRepositories, releases
}
