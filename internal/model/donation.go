package model

import (
	"time"
)

type CampaignID = int32

type UserDonation struct {
	ID         int32
	CampaignID CampaignID
	Amount     int32
	DonatedAt  time.Time
}

type Campaign struct {
	ID          CampaignID
	OrgID       OrgID
	Title       string
	Description string
	GoalAmount  int32
	TotalAmount int64
}

type CreateCampaignParams struct {
	OrganizationID OrgID
	Title          string
	Description    string
	GoalAmount     int32
	StartDate      time.Time
	EndDate        time.Time
}
