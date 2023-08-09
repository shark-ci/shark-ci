package store

import (
	"context"
	"time"

	"log/slog"

	"github.com/google/uuid"
	"github.com/shark-ci/shark-ci/ci-server/models"
	"github.com/shark-ci/shark-ci/ci-server/types"
	"golang.org/x/oauth2"
)

type Storer interface {
	Ping(ctx context.Context) error
	Close(ctx context.Context) error

	// Cleanup expired OAuth2 states.
	Clean(ctx context.Context) error

	GetAndDeleteOAuth2State(ctx context.Context, state uuid.UUID) (*models.OAuth2State, error)
	CreateOAuth2State(ctx context.Context, state *models.OAuth2State) error

	GetUser(ctx context.Context, userID int64) (*models.User, error)
	CreateUserAndServiceUser(ctx context.Context, serviceUser *models.ServiceUser) (int64, int64, error)

	GetServiceUserIDsByServiceUsername(ctx context.Context, service string, username string) (int64, int64, error)
	GetServiceUsersRepoFetchInfo(ctx context.Context, userID int64) ([]*types.ServiceUserRepoFetchInfo, error)
	UpdateServiceUserToken(ctx context.Context, serviceUserID int64, token *oauth2.Token) error

	GetRepoIDByServiceRepoID(ctx context.Context, service string, serviceRepoID int64) (int64, error)
	GetReposByUser(ctx context.Context, userID int64) ([]models.Repo, error)
	GetRepoWebhookChangeInfo(ctx context.Context, repoID int64) (*types.RepoWebhookChangeInfo, error)
	CreateOrUpdateRepos(ctx context.Context, repos []models.Repo) error
	UpdateRepoWebhook(ctx context.Context, repoID int64, webhookID *int64) error

	//GetPipeline(ctx context.Context, pipelineID int64) (*models.Pipeline, error)
	GetPipelineCreationInfo(ctx context.Context, repoID int64) (*types.PipelineCreationInfo, error)
	CreatePipeline(ctx context.Context, pipeline *models.Pipeline) (int64, error)
	UpdatePipelineStatus(ctx context.Context, pipelineID int64, status string, started_at *time.Time, finished_at *time.Time) error

	GetPipelineStateChangeInfo(ctx context.Context, pipelineID int64) (*types.PipelineStateChangeInfo, error)
}

func Cleaner(s Storer, d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for {
			<-ticker.C
			err := s.Clean(context.TODO())
			if err != nil {
				slog.Warn("store: databse cleanup failed", "err", err)
			}
		}
	}()
}
