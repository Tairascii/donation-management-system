package handler

import (
	"net/http"

	"github.com/Tairascii/donation-managment-system/pkg/response_writers"
)

func (h *Handler) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	id, err := CampaignIDFromPath(r)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteCampaign(ctx, id)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
