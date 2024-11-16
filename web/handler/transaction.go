package handler

import (
	"go_startup_api/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) Index(c *gin.Context) {
	transaction, err := h.TransactionService.GetAllTransactions()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transaction})
}
