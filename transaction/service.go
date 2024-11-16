package transaction

import (
	"errors"
	"go_startup_api/campaign"
	"go_startup_api/payment"
	"strconv"
)

type Service interface {
	GetTransactionByCampaignID(input GetTransactionCampaignInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
	GetAllTransactions() ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionByCampaignID(input GetTransactionCampaignInput) ([]Transaction, error) {
	campaignID := input.ID

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of this campaign")
	}

	transactions, err := s.repository.GetByCampaignID(campaignID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.UserID = input.User.ID

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	transaction.PaymentURL = paymentUrl
	updateTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return updateTransaction, err
	}

	return updateTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transactionID, _ := strconv.Atoi(input.OrderID)
	transaction, err := s.repository.GetByID(transactionID)
	if err != nil {
		return err
	}

	if transaction.Status != "pending" {
		return errors.New("transaction is not pending")
	}

	if (input.PaymentType == "credit_card") && (input.TransactionStatus == "capture") && (input.FraudStatus == "accept") {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if (input.TransactionStatus == "deny") || (input.TransactionStatus == "expire") || (input.TransactionStatus == "cancel") {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
		if err != nil {
			return err
		}

		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount
		campaign.BackerCount = campaign.BackerCount + 1
		_, err = s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetAllTransactions() ([]Transaction, error) {
	transactions, err := s.repository.GetAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
