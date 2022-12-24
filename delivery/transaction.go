package delivery

import (
	"jangFundraising/helper"
	"jangFundraising/transaction"
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

	transactions, err := h.usecase.GetTransactionsByCampaignID(input)
	if err != nil {
		resp := helper.APIResponse("failed getting campaign transactions detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("successfully retrieved transactions detail", http.StatusOK, "success", transactions)
	c.JSON(http.StatusBadRequest, resp)
}
