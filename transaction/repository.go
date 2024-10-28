package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
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
func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	transaction := []Transaction{}

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = ?", 1).Where("user_id = ?", userID).Order("id DESC").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
