package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Tairascii/donation-managment-system/internal/model"
	"github.com/Tairascii/donation-managment-system/pkg/response_writers"
)

type CreateCampaignPayload struct {
	OrgID       int32  `json:"org_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	GoalAmount  int32  `json:"goal_amount"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type CreateCampaignResponse struct {
	ID int32 `json:"id"`
}

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	var payload CreateCampaignPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse(time.DateOnly, payload.StartDate)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse(time.DateOnly, payload.EndDate)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.uc.CreateCampaign(ctx, model.CreateCampaignParams{
		OrganizationID: payload.OrgID,
		Title:          payload.Title,
		Description:    payload.Description,
		GoalAmount:     payload.GoalAmount,
		StartDate:      startDate,
		EndDate:        endDate,
	})
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response_writers.JSONResponseWriter(w, http.StatusOK, CreateCampaignResponse{ID: id})
}
