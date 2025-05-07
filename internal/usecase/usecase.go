package usecase

import (
	"context"

	"github.com/Tairascii/donation-managment-system/internal/model"
)

type Repo interface {
	Donate(ctx context.Context, userID model.UserID, campaignID model.CampaignID, amount int32) error
	DonationsByUserID(ctx context.Context, userID model.UserID) ([]model.UserDonation, error)
	CampaignByID(ctx context.Context, id model.CampaignID) (model.Campaign, error)
	DeleteCampaign(ctx context.Context, id model.CampaignID) error
	CreateCampaign(ctx context.Context, p model.CreateCampaignParams) (model.CampaignID, error)
}

type UseCase struct {
	repo Repo
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) Donate(ctx context.Context, userID model.UserID, campaignID model.CampaignID, amount int32) error {
	return u.repo.Donate(ctx, userID, campaignID, amount)
}

func (u *UseCase) DonationsByUserID(ctx context.Context, userID model.UserID) ([]model.UserDonation, error) {
	return u.repo.DonationsByUserID(ctx, userID)
}

func (u *UseCase) CampaignByID(ctx context.Context, id model.CampaignID) (model.Campaign, error) {
	return u.repo.CampaignByID(ctx, id)
}

func (u *UseCase) DeleteCampaign(ctx context.Context, id model.CampaignID) error {
	return u.repo.DeleteCampaign(ctx, id)
}

func (u *UseCase) CreateCampaign(ctx context.Context, p model.CreateCampaignParams) (model.CampaignID, error) {
	return u.repo.CreateCampaign(ctx, p)
}
