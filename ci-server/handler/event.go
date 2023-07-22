package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"

	ciserver "github.com/FilipSolich/shark-ci/ci-server"
	"github.com/FilipSolich/shark-ci/ci-server/service"
	"github.com/FilipSolich/shark-ci/ci-server/store"
	"github.com/FilipSolich/shark-ci/shared/message_queue"
)

type EventHandler struct {
	l        *slog.Logger
	s        store.Storer
	mq       message_queue.MessageQueuer
	services service.Services
}

func NewEventHandler(l *slog.Logger, s store.Storer, mq message_queue.MessageQueuer, services service.Services) *EventHandler {
	return &EventHandler{
		l:        l,
		s:        s,
		mq:       mq,
		services: services,
	}
}

func (h *EventHandler) HandleEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	serviceName, ok := params["service"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv, ok := h.services[serviceName]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	job, err := srv.HandleEvent(r)
	if err != nil {
		if errors.Is(err, service.NoErrPingEvent) {
			w.Write([]byte("pong"))
		} else if errors.Is(err, service.ErrEventNotSupported) {
			http.Error(w, "cannot handle this type of event", http.StatusNotImplemented)
		} else {
			h.l.Error("service: cannot hadle event", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = h.s.CreateJob(context.TODO(), job)
	if err != nil {
		h.l.Error("store: cannot create job", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.mq.SendJob(context.TODO(), job)
	if err != nil {
		h.l.Error("store: cannot send job", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	repo, err := h.s.GetRepo(ctx, job.RepoID)
	if err != nil {
		h.l.Error("store: cannot get repo", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serviceUser, err := h.s.GetServiceUserByRepo(ctx, repo)
	if err != nil {
		h.l.Error("store: cannot get service user", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status := service.NewStatus(service.StatusPending, job.TargetURL, ciserver.CIServer, "Job in progress")
	err = srv.CreateStatus(ctx, serviceUser, repo, job, status)
	if err != nil {
		h.l.Error("cannot create status", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}