package models

type Repo struct {
	ID            int64  `json:"id"`
	Service       string `json:"service"`
	Owner         string `json:"owner"`
	Name          string `json:"name"`
	RepoServiceID int64  `json:"repo_service_id"`
	WebhookID     *int64 `json:"webhook_id,omitempty"`
	ServiceUserID int64  `json:"service_user_id"`
}
