package handler

import (
	"context"
	"net/http"

	"github.com/Tairascii/donation-managment-system/internal/model"
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
	return router
}
