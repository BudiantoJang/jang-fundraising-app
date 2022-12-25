package transaction

import (
	"errors"
	"jangFundraising/campaign"
	"jangFundraising/payment"
)

type Usecase interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type usecase struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentUsecase     payment.Usecase
}

func NewUsecase(repository Repository, campaignRepository campaign.Repository, paymentUsecase payment.Usecase) *usecase {
	return &usecase{repository, campaignRepository, paymentUsecase}
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

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentUrl, err := u.paymentUsecase.GetPaymentUrl(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentUrl

	newTransaction, err = u.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil

}
