package main

import (
	"github.com/arldka/github-releases-gchat-notifier/internal/scraper"
	"github.com/arldka/github-releases-gchat-notifier/internal/storage"
)

func main() {

	db := storage.NewDB()
	gh := scraper.NewClient()

	// Retrieve the list of repositories
	repositories := db.ListRepositories()

	// Read the releases of all the concerned repositories
	repositories, releases := scraper.GetReleases(repositories)

	// Remove all the repositories that already generated notifications

	// Send notifications for the remaining repositories repositories
}
