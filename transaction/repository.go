package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("campaign_id = ?", campaignID).Preload("User").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("user_id = ?", userID).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {

	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
