package transaction

import (
	"jangFundraising/campaign"
	"jangFundraising/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
	Campaign   campaign.Campaign
}

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	CampaignID int `json:"campaignId" binding:"required"`
	Amount     int `json:"amount" binding:"required"`
	User       user.User
}
