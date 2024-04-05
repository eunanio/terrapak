package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"terrapak/internal/api/webhook/repo"
	"terrapak/internal/api/webhook/rest"
	"text/template"
)

type CommentResponse struct {
	ID int `json:"id"`
}

const REPORT_TEMPLATE = `## New Module Changes Detected
Terrapak has detected changes in the following modules:
| Module | Version |
| :---: | :---: |{{ range . }}
| {{.Name}} | {{.Version}} |{{ end }}
`

func CreateSyncReport(event PullRequestEvent, syncReport SyncReport) error {
	if len(syncReport.Items) > 0 {
		var mdTable bytes.Buffer
		tmeplate, err := template.New("report").Parse(REPORT_TEMPLATE); if err != nil {
			return err
		}

		err = tmeplate.Execute(&mdTable, syncReport.Items); if err != nil {
			return err
		}
		err = createComment(event, mdTable.String()); if err != nil {
			slog.Error(err.Error())
			return err
		}
	}

	return nil
}

func createComment(event PullRequestEvent, markdown string) error {
	access_token, err := repo.GetAccessToken(event.Installation.ID); if err != nil {
		return err
	}
	req := CommentBodyRequest{Body: markdown}
	client := rest.New(access_token)
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/comments", event.PullRequest.Head.Repo.FullName, event.PullRequest.Number)
	jsonData, err := json.Marshal(req); if err != nil {
		return err
	}

	res, err := client.Post(endpoint, "application/vnd.github.raw+json",bytes.NewBuffer(jsonData)); if err != nil {
		slog.Error(err.Error())
		return err
	}

	if res.StatusCode != 201 {
		return fmt.Errorf("failed to create comment: %s", res.Status)
	}

	return nil
}
