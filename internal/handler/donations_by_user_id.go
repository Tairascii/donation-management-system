package handler

import (
	"net/http"
	"time"

	"github.com/Tairascii/donation-managment-system/internal/model"
	"github.com/Tairascii/donation-managment-system/pkg/response_writers"
)

type Donation struct {
	ID         int32            `json:"id"`
	CampaignID model.CampaignID `json:"campaign_id"`
	Amount     int32            `json:"amount"`
	DonatedAt  time.Time        `json:"donated_at"`
}

type DonationsByUserIDResponse struct {
	Donations []Donation `json:"donations"`
}

func (h *Handler) DonationsByUserID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	userID, err := UserIDFromPath(r)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	donations, err := h.uc.DonationsByUserID(ctx, userID)
	if err != nil {
		response_writers.ErrorResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := make([]Donation, len(donations))
	for i, donation := range donations {
		res[i] = Donation{
			ID:         donation.ID,
			CampaignID: donation.CampaignID,
			Amount:     donation.Amount,
			DonatedAt:  donation.DonatedAt,
		}
	}

	response_writers.JSONResponseWriter(w, http.StatusOK, DonationsByUserIDResponse{Donations: res})
}
