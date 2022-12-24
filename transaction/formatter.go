package transaction

import (
	"time"
)

type CampaignTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserTransactionsFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"createdAt"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionsFormatter {
	formatted := CampaignTransactionsFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatted
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionsFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionsFormatter{}
	}

	var formattedTransaction []CampaignTransactionsFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		formattedTransaction = append(formattedTransaction, formatter)
	}

	return formattedTransaction
}

func FormatUserTransaction(transaction Transaction) UserTransactionsFormatter {
	formatted := UserTransactionsFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	campaignFormatted := CampaignFormatter{
		Name:     transaction.Campaign.Name,
		ImageUrl: "",
	}

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatted.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatted.Campaign = campaignFormatted

	return formatted
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionsFormatter {
	if len(transactions) == 0 {
		return []UserTransactionsFormatter{}
	}

	var formattedTransaction []UserTransactionsFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		formattedTransaction = append(formattedTransaction, formatter)
	}

	return formattedTransaction
}
