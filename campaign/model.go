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
	GoalAmount     int
	CurrentAmount  int
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

type CreateCampaignInput struct {
	Name        string `json:"name" binding:"required"`
	Summary     string `json:"summary" binding:"required"`
	Description string `json:"description" binding:"required"`
	GoalAmount  int    `json:"goalAmount" binding:"required"`
	Perks       string `json:"perks" binding:"required"`
	User        user.User
}
