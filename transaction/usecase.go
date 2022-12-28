package transaction

import (
	"errors"
	"fmt"
	"jangFundraising/campaign"
	"jangFundraising/payment"
	"strconv"
)

type Usecase interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	GetPaymentProcess(input TransactionNotificationInput) error
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

func (u *usecase) GetPaymentProcess(input TransactionNotificationInput) error {
	fmt.Println(input)
	transactonID, _ := strconv.Atoi(input.OrderID)
	fmt.Println(transactonID)
	trans, err := u.repository.GetByID(transactonID)
	if err != nil {
		return err
	}

	fmt.Println(trans)
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		trans.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		trans.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		trans.Status = "cancelled"
	}
	fmt.Println(trans.Status)

	updatedTrans, err := u.repository.Update(trans)
	if err != nil {
		return err
	}
	fmt.Println(updatedTrans)

	campaign, err := u.campaignRepository.FindByID(updatedTrans.CampaignID)
	if err != nil {
		return err
	}
	fmt.Println(campaign)

	if updatedTrans.Status == "paid" {
		campaign.DonatorCount = campaign.DonatorCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTrans.Amount
		fmt.Println(campaign)

		_, err := u.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
