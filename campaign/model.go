package campaign

import (
	"jangFundraising/user"
	"time"
)

type Campaign struct {
	ID             int
	UserID         int
	Name           string
	Summary        string
	Description    string
	Perks          string
	DonatorCount   int
	GoalAmmount    int
	CurrentAmmount int
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImage
	User           user.User
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
