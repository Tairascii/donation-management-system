package repository

import (
	"context"

	"github.com/Tairascii/donation-managment-system/db/query"
	"github.com/Tairascii/donation-managment-system/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) DonationsByUserID(ctx context.Context, userID model.UserID) ([]model.UserDonation, error) {
	q := query.New(r.db)
	res, err := q.DonationsByUserID(ctx, pgtype.Int4{Int32: userID, Valid: true})
	if err != nil {
		return nil, err
	}

	donations := make([]model.UserDonation, len(res))
	for i, d := range res {
		donations[i] = model.UserDonation{
			ID:         d.ID,
			CampaignID: d.CampaignID.Int32,
			Amount:     d.Amount,
			DonatedAt:  d.DonatedAt.Time,
		}
	}
	return donations, nil
}
