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
		resp := helper.APIResponse("failed getting 1 transactions detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User.ID = currentUser.ID

	transactions, err := h.usecase.GetTransactionsByCampaignID(input)
	if err != nil {
		resp := helper.APIResponse("failed getting 2 transactions detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("successfully retrieved 3 detail", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, resp)
}
