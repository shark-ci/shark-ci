package store

import (
	"context"
	"time"

	"github.com/FilipSolich/shark-ci/ci-server/models"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
)

type Storer interface {
	Ping(ctx context.Context) error
	Close(ctx context.Context) error

	// Cleanup expired OAuth2 states.
	Clean(ctx context.Context) error

	GetOAuth2State(ctx context.Context, state uuid.UUID) (*models.OAuth2State, error)
	CreateOAuth2State(ctx context.Context, state *models.OAuth2State) error
	DeleteOAuth2State(ctx context.Context, state *models.OAuth2State) error

	GetUser(ctx context.Context, id int64) (*models.User, error)
	CreateUserAndServiceUser(ctx context.Context, serviceUser *models.ServiceUser) (int64, error)

	GetServiceUserByUniqueName(ctx context.Context, service string, username string) (*models.ServiceUser, error)
	GetServiceUserByRepo(ctx context.Context, repoID int64) (*models.ServiceUser, error)
	GetServiceUserByUserAndService(ctx context.Context, userID int64, service string) (*models.ServiceUser, error)
	GetServiceUsersByUser(ctx context.Context, userID int64) ([]models.ServiceUser, error)
	UpdateServiceUserToken(ctx context.Context, serviceUser *models.ServiceUser, token *oauth2.Token) error

	GetRepo(ctx context.Context, repoID int64) (*models.Repo, error)
	GetRepoName(ctx context.Context, repoID int64) (string, error)
	GetRepoIDByServiceRepoID(ctx context.Context, service string, serviceRepoID int64) (int64, error)
	GetReposByUser(ctx context.Context, userID int64) ([]models.Repo, error)
	CreateOrUpdateRepos(ctx context.Context, repos []models.Repo) error
	UpdateRepoWebhook(ctx context.Context, repoID int64, webhookID *int64) error

	GetPipeline(ctx context.Context, pipelineID int64) (*models.Pipeline, error)
	CreatePipeline(ctx context.Context, pipeline *models.Pipeline) error
	UpdatePipelineStatus(ctx context.Context, pipelineID int64, status string, started_at *time.Time, finished_at *time.Time) error
}

func Cleaner(s Storer, d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for {
			<-ticker.C
			err := s.Clean(context.TODO())
			if err != nil {
				slog.Error("store: databse cleanup failed", "err", err)
			}
		}
	}()
}
