package repository

import (
	"context"

	"github.com/Tairascii/donation-managment-system/db/query"
	"github.com/Tairascii/donation-managment-system/internal/model"
)

func (r *Repository) CampaignByID(ctx context.Context, id model.CampaignID) (model.Campaign, error) {
	q := query.New(r.db)
	res, err := q.CampaignByID(ctx, id)
	if err != nil {
		return model.Campaign{}, err
	}

	return model.Campaign{
		ID:          res.ID,
		OrgID:       res.OrganizationID.Int32,
		Title:       res.Title,
		Description: res.Description.String,
		GoalAmount:  res.GoalAmount.Int32,
		TotalAmount: res.TotalAmount.Int64,
	}, nil
}
