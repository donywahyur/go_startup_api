package transaction

import (
	"fmt"
	"go_startup_api/campaign"
)

type Service interface {
	GetTransactionByCampaignID(input GetTransactionCampaignInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetTransactionCampaignInput) ([]Transaction, error) {
	campaignID := input.ID

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, fmt.Errorf("not an owner of this campaign")
	}

	transactions, err := s.repository.GetByCampaignID(campaignID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
