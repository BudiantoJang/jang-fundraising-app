package delivery

import (
	"jangFundraising/helper"
	"jangFundraising/transaction"
	"jangFundraising/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	usecase transaction.Usecase
}

func NewTransactionHandler(service transaction.Usecase) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		resp := helper.APIResponse("failed getting campaign transactions detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User.ID = currentUser.ID

	transactions, err := h.usecase.GetTransactionsByCampaignID(input)
	if err != nil {
		resp := helper.APIResponse("failed getting campaign transactions detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("successfully retrieved campaign detail", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, resp)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.usecase.GetTransactionsByUserID(userID)
	if err != nil {
		resp := helper.APIResponse("failed getting user transactions detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("successfully retrieved user transactions detail", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, resp)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	currentUser := c.MustGet("currentUser").(user.User)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		resp := helper.APIResponse("failed binding payload", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	input.User = currentUser

	newTransaction, err := h.usecase.CreateTransaction(input)
	if err != nil {
		resp := helper.APIResponse("failed creating new transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("successfully created new transaction", http.StatusOK, "success", newTransaction)
	c.JSON(http.StatusOK, resp)
}
