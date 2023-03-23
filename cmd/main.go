package main

import (
	"fmt"

	"github.com/arldka/github-releases-gchat-notifier/internal/notifications"
	"github.com/arldka/github-releases-gchat-notifier/internal/scraper"
	"github.com/arldka/github-releases-gchat-notifier/internal/storage"
)

func main() {

	fmt.Println("Before NewDB")
	db := storage.NewDB()
	fmt.Println("After NewDB and before NewClient")
	gh := scraper.NewClient()
	fmt.Println("After NewClient")

	fmt.Println("Before db.ListRepositories")
	// Retrieve the list of repositories
	repositories := db.ListRepositories()

	fmt.Println("oui")
	fmt.Println(repositories)
	fmt.Println("non")
	// Read the releases of all the concerned repositories
	repositories, releases := gh.GetReleases(repositories)

	// Remove all the repositories that already generated notifications
	db.UpdateRepositories(repositories)

	fmt.Println(releases)

	// Send Notifications
	notifications.SendNotifications(releases)

	// Mark all repositories as notified
}
