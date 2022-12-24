package transaction

import (
	"errors"
	"jangFundraising/campaign"
)

type Usecase interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

type usecase struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewUsecase(repository Repository, campaignRepository campaign.Repository) *usecase {
	return &usecase{repository, campaignRepository}
}

func (u *usecase) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := u.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.User.ID != input.User.ID {
		return []Transaction{}, errors.New("not authorized")
	}

	transactions, err := u.repository.FindByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
