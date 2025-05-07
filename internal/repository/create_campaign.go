package repository

import (
	"context"

	"github.com/Tairascii/donation-managment-system/db/query"
	"github.com/Tairascii/donation-managment-system/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) CreateCampaign(ctx context.Context, p model.CreateCampaignParams) (model.CampaignID, error) {
	q := query.New(r.db)
	return q.CreateCampaign(ctx, &query.CreateCampaignParams{
		OrganizationID: pgtype.Int4{Int32: p.OrganizationID, Valid: true},
		Title:          p.Title,
		Description:    pgtype.Text{String: p.Description, Valid: true},
		GoalAmount:     pgtype.Int4{Int32: p.GoalAmount, Valid: true},
		StartDate:      pgtype.Date{Time: p.StartDate, Valid: true},
		EndDate:        pgtype.Date{Time: p.EndDate, Valid: true},
	})
}
