package mine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Story struct {
	ID   int    `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

// Command returns a list of tickets assigned to the current user.
var Command = &cobra.Command{
	Use:   "mine",
	Short: "List my tickets",
	Long:  "List tickets assigned to my account",
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Root().PersistentFlags().GetString("project-id")
		username, _ := cmd.Root().PersistentFlags().GetString("username")
		token := os.Getenv("PIVOTAL_TOKEN")

		url := fmt.Sprintf(`https://pivotaltracker.com/services/v5/projects/%s/stories?filter=mywork:"%s"`, projectID, username)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Add("X-TrackerToken", token)

		httpClient := &http.Client{
			Timeout: time.Second * 10,
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Printf("Failed to list stories: %v\n", err)
			return
		}
		defer resp.Body.Close()

		var stories []Story
		if err := json.NewDecoder(resp.Body).Decode(&stories); err != nil {
			fmt.Printf("Failed to decode API response: %v\n", err)
			return
		}

		fmt.Printf("Found %d stories:\n", len(stories))
		for _, s := range stories {
			fmt.Printf("- %s (#%d) -> %s\n", s.Name, s.ID, s.URL)
		}
	},
}
