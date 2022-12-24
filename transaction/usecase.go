package transaction

type Usecase interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	transactions, err := u.repository.FindByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
