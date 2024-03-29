package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/shark-ci/shark-ci/internal/server/middleware"
	"github.com/shark-ci/shark-ci/internal/server/service"
	"github.com/shark-ci/shark-ci/internal/server/store"
	"github.com/shark-ci/shark-ci/internal/server/types"
	"github.com/shark-ci/shark-ci/templates"
)

type RepoHandler struct {
	s        store.Storer
	services service.Services
}

func NewRepoHandler(s store.Storer, services service.Services) *RepoHandler {
	return &RepoHandler{
		s:        s,
		services: services,
	}
}

func (h *RepoHandler) FetchUnregistredRepos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.UserFromContext(ctx, w)

	srvName := mux.Vars(r)["service"]
	srv, ok := h.services[srvName]
	if !ok {
		Error400(w, fmt.Sprintf("Unknown service %s", srvName))
		return
	}

	serviceUser, err := h.s.GetServiceUserByUserID(ctx, srv.Name(), user.ID)
	if err != nil {
		Error5xx(w, http.StatusInternalServerError, "Cannot get service user.", err)
		return
	}

	repos, err := srv.GetUserRepos(ctx, serviceUser.Token(), serviceUser.ID)
	if err != nil {
		Error5xx(w, http.StatusInternalServerError, "Cannot get user unregistered repos.", err)
		return
	}

	err = templates.ReposRegisterTmpl.Execute(w, map[string]any{
		"Repos":          repos,
		csrf.TemplateTag: csrf.TemplateField(r),
	})
	if err != nil {
		Error5xx(w, http.StatusInternalServerError, "Cannot execute template.", err)
		return
	}
}

func (h *RepoHandler) HandleRegisterRepo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.UserFromContext(ctx, w)

	err := r.ParseForm()
	if err != nil {
		Error400(w, "Invalid data")
	}

	srvName := r.FormValue("service")
	repoIDString := r.FormValue("repo_id")
	owner := r.FormValue("owner")
	repoName := r.FormValue("name")

	srv, ok := h.services[srvName]
	if !ok {
		Error400(w, fmt.Sprintf("Unknown service %s", srvName))
		return
	}

	repoID, err := strconv.ParseInt(repoIDString, 10, 64)
	if err != nil {
		Error400(w, "Unknown repo")
		return
	}

	serviceUser, err := h.s.GetServiceUserByUserID(ctx, srv.Name(), user.ID)
	if err != nil {
		Error5xx(w, http.StatusInternalServerError, "Cannot get service user.", err)
		return
	}

	hookID, err := srv.CreateWebhook(ctx, serviceUser.Token(), owner, repoName)
	if err != nil {
		Error5xx(w, http.StatusInternalServerError, "Cannot create webhook", err)
		return
	}

	_, err = h.s.CreateRepo(ctx, types.Repo{
		Service:       srv.Name(),
		Owner:         owner,
		Name:          repoName,
		RepoServiceID: repoID, // TODO: repoServiceID but ServiceUserID. Change name of one of them.
		WebhookID:     hookID,
		ServiceUserID: serviceUser.ID,
	})
	if err != nil {
		Error5xx(w, http.StatusInternalServerError, "Cannot create repo", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

//func (h *RepoHandler) HandleUnregisterRepo(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	serviceUser, repo, srv, err := h.getInfoFromRequest(ctx, w, r)
//	if err != nil {
//		slog.Error("cannot get info from request", "err", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	err = srv.DeleteWebhook(ctx, serviceUser, repo)
//	if err != nil {
//		slog.Error("service: cannot delete webhook", "err", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	err = h.s.UpdateRepoWebhook(ctx, repo)
//	if err != nil {
//		slog.Error("store: cannot update webhook", "err", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	http.Redirect(w, r, "/repositories", http.StatusFound)
//}
//
//func (h *RepoHandler) HandleActivateRepo(w http.ResponseWriter, r *http.Request) {
//	h.changeRepoState(w, r, true)
//}
//
//func (h *RepoHandler) HandleDeactivateRepo(w http.ResponseWriter, r *http.Request) {
//	h.changeRepoState(w, r, false)
//}
//
//func (h *RepoHandler) changeRepoState(w http.ResponseWriter, r *http.Request, active bool) {
//	ctx := r.Context()
//	serviceUser, repo, srv, err := h.getInfoFromRequest(ctx, w, r)
//	if err != nil {
//		slog.Error("cannot get infor from request", "err", err)
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	repo, err = srv.ChangeWebhookState(ctx, serviceUser, repo, active)
//	if err != nil {
//		slog.Error("service: cannot change a webhook state", "err", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	err = h.s.UpdateRepoWebhook(ctx, repo)
//	if err != nil {
//		slog.Error("store: cannot update a webhook", "err", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	http.Redirect(w, r, "/repositories", http.StatusFound)
//}

//func (h *RepoHandler) getInfoFromRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) (*models.ServiceUser, *models.Repo, service.ServiceManager, error) {
//	user, ok := middleware.UserFromContext(ctx, w)
//	if !ok {
//		return nil, nil, nil, errors.New("unauthorized user")
//	}
//
//	r.ParseForm()
//	repo, err := h.s.GetRepo(ctx, r.Form.Get("repo_id"))
//	if err != nil {
//		return nil, nil, nil, err
//	}
//
//	srv, ok := h.services[repo.ServiceName]
//	if !ok {
//		return nil, nil, nil, fmt.Errorf("unknown service: %s", repo.ServiceName)
//	}
//
//	serviceUser, err := h.s.GetServiceUserByUser(ctx, user, repo.ServiceName)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//
//	return serviceUser, repo, srv, nil
//}
//
//func splitRepos(repos []*models.Repo) ([]*models.Repo, []*models.Repo) {
//	registered := []*models.Repo{}
//	notRegistered := []*models.Repo{}
//	for _, repo := range repos {
//		if repo.WebhookID == 0 {
//			notRegistered = append(notRegistered, repo)
//		} else {
//			registered = append(registered, repo)
//		}
//	}
//	return registered, notRegistered
//}
