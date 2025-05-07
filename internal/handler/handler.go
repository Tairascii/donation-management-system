package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Tairascii/donation-managment-system/internal/model"
)

const (
	campaignIDPathParam = "campaign_id"
	userIDPathParam     = "user_id"
)

var (
	ErrInvalidCampaignID = errors.New("invalid campaign id")
	ErrInvalidUserID     = errors.New("invalid user id")
)

type UseCase interface {
	Donate(ctx context.Context, userID model.UserID, campaignID model.CampaignID, amount int32) error
	DonationsByUserID(ctx context.Context, userID model.UserID) ([]model.UserDonation, error)
	CampaignByID(ctx context.Context, id model.CampaignID) (model.Campaign, error)
	DeleteCampaign(ctx context.Context, id model.CampaignID) error
	CreateCampaign(ctx context.Context, p model.CreateCampaignParams) (model.CampaignID, error)
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
	router.HandleFunc("POST /campaign", h.CreateCampaign)
	router.HandleFunc("GET /campaign/{campaign_id}", h.CampaignByID)
	router.HandleFunc("DELETE /campaign/{campaign_id}", h.DeleteCampaign)
	router.HandleFunc("POST /campaign/{campaign_id}/donate", h.Donate)
	router.HandleFunc("GET /user/{user_id}/donation", h.DonationsByUserID)
	return router
}

func CampaignIDFromPath(r *http.Request) (model.CampaignID, error) {
	campaignIDRaw := r.PathValue(campaignIDPathParam)
	campaignID, err := strconv.Atoi(campaignIDRaw)
	if err != nil || campaignID <= 0 {
		return 0, ErrInvalidCampaignID
	}
	return model.CampaignID(campaignID), nil
}

func UserIDFromPath(r *http.Request) (model.CampaignID, error) {
	userIDRaw := r.PathValue(userIDPathParam)
	userID, err := strconv.Atoi(userIDRaw)
	if err != nil || userID <= 0 {
		return 0, ErrInvalidUserID
	}
	return model.CampaignID(userID), nil
}
