package notifications

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/chat/v1"

	"github.com/arldka/github-releases-gchat-notifier/internal/models"
)

func SendNotifications(releases *[]models.Release) {
	// Set up the message to be sent.

	for _, release := range *releases {

		parts := strings.Split(release.RepoName, "/")
		if len(parts) != 2 {
			log.Fatalf("Invalid repository name: %s", release.RepoName)
		}
		owner := parts[0]

		message := &chat.Message{
			Cards: []*chat.Card{
				{
					Header: &chat.CardHeader{
						Title:      release.RepoName,
						Subtitle:   fmt.Sprintf("New version released %s", release.Tag),
						ImageUrl:   fmt.Sprintf("https://avatars.githubusercontent.com/%s", owner),
						ImageStyle: "IMAGE",
					},
					Sections: []*chat.Section{
						{
							Widgets: []*chat.WidgetMarkup{
								{
									Buttons: []*chat.Button{
										{
											TextButton: &chat.TextButton{
												Text: "View on Github",
												OnClick: &chat.OnClick{
													OpenLink: &chat.OpenLink{
														Url: release.ReleaseURL,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		// Convert the message to JSON.
		jsonMessage, err := message.MarshalJSON()
		if err != nil {
			log.Fatalf("Failed to marshal message: %v", err)
		}

		// Set up the webhook URL and secret.
		webhookURL := os.Getenv("WEBHOOK_URL")

		// Create an HTTP request to send the message.
		req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonMessage))
		if err != nil {
			log.Fatalf("Failed to create HTTP request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send the message via the webhook.
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		defer resp.Body.Close()

		log.Println("Message sent successfully!")
	}

}
