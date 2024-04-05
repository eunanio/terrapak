package webhook

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func HandleGithubWebhook(c *gin.Context) {
	event := PullRequestEvent{}
	err := c.Bind(&event); if err != nil {
		slog.Error(err.Error())
	}
	// eventMessage, _ := json.Marshal(event)
	// os.WriteFile("event.json",eventMessage, 0644)
	// fmt.Println(event.Action)
	switch event.Action {
		case "opened":
			onSyncPullRequest(event)
		case "closed":
			onClosePullRequest(event)
		case "synchronize":
			onSyncPullRequest(event)
		case "reopened":
			onSyncPullRequest(event)
	}

	c.JSON(200, gin.H{})
}