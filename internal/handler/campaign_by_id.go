package handler

import (
	"net/http"

	"github.com/Tairascii/donation-managment-system/internal/model"
	"github.com/Tairascii/donation-managment-system/pkg/response_writers"
)

type CampaignByIDResponse struct {
	ID          model.CampaignID `json:"id"`
	OrgID       model.OrgID      `json:"org_id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	GoalAmount  int32            `json:"goal_amount"`
	TotalAmount int64            `json:"total_amount"`
}

func (h *Handler) CampaignByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	id, err := CampaignIDFromPath(r)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	campaign, err := h.uc.CampaignByID(ctx, id)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response_writers.JSONResponseWriter(w, http.StatusOK, CampaignByIDResponse{
		ID:          campaign.ID,
		OrgID:       campaign.OrgID,
		Title:       campaign.Title,
		Description: campaign.Description,
		GoalAmount:  campaign.GoalAmount,
		TotalAmount: campaign.TotalAmount,
	})
}
