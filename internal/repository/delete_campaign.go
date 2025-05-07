package repository

import (
	"context"

	"github.com/Tairascii/donation-managment-system/db/query"
	"github.com/Tairascii/donation-managment-system/internal/model"
)

func (r *Repository) DeleteCampaign(ctx context.Context, id model.CampaignID) error {
	q := query.New(r.db)
	return q.DeleteCampaign(ctx, id)
}
