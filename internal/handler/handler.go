package handler

import (
	"net/http"
)

type UseCase interface {
}

type Handler struct {
	uc UseCase
}

func New(uc UseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func AttachRoutes(h *Handler) http.Handler {
	router := http.NewServeMux()
	return router
}
