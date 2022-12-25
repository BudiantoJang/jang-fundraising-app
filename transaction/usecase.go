package transaction

import (
	"errors"
	"jangFundraising/campaign"
)

type Usecase interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
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

func (u *usecase) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := u.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (u *usecase) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		CampaignID: input.CampaignID,
		Amount:     input.Amount,
		UserID:     input.User.ID,
		Status:     "pending",
	}

	newTransaction, err := u.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil

}
