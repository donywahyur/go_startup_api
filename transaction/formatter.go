package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	transactionFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	transactionFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.CampaignID = transaction.CampaignID
	formatter.UserID = transaction.UserID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL

	return formatter
}
