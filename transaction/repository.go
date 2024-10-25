package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	transaction := []Transaction{}

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id DESC").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
