package repository

import (
	"context"

	"github.com/Tairascii/donation-managment-system/db/query"
	"github.com/Tairascii/donation-managment-system/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) Donate(ctx context.Context, userID model.UserID, campaignID model.CampaignID, amount int32) error {
	q := query.New(r.db)
	return q.Donate(ctx, &query.DonateParams{
		UserID:     pgtype.Int4{Int32: userID, Valid: true},
		CampaignID: pgtype.Int4{Int32: campaignID, Valid: true},
		Amount:     amount,
	})
}
