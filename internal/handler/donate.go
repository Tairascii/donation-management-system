package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tairascii/donation-managment-system/pkg/response_writers"
)

type DonatePayload struct {
	UserID int32 `json:"user_id"`
	Amount int32 `json:"amount"`
}

func (h *Handler) Donate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	campaignID, err := CampaignIDFromPath(r)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload DonatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.Donate(ctx, payload.UserID, campaignID, payload.Amount)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
