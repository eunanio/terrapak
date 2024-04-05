package webhook

type PullRequestEvent struct {
	Action       string       `json:"action"`
	Number       int          `json:"number"`
	PullRequest  PullRequest  `json:"pull_request"`
	Installation Installation `json:"installation"`
}

type PullRequest struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
	Number int    `json:"number"`
	State  string `json:"state"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Merged bool   `json:"merged"`
	Head   GitRef `json:"head"`
	Base   GitRef `json:"base"`
}

type Installation struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
}

type User struct {
	Login string `json:"login"`
}

type GitRef struct {
	Ref  string     `json:"ref"`
	Sha  string     `json:"sha"`
	Repo Repository `json:"repo"`
}

type Repository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	CloneURL string `json:"clone_url"`
}

type SyncReport struct {
	Items []SyncReportItem `json:"items"`
}

type SyncReportItem struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CommentBodyRequest struct {
	Body string `json:"body"`
}
