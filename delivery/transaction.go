package delivery

import (
	"fmt"
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
	fmt.Println(transactions)
	fmt.Println("AAAAA")
	resp := helper.APIResponse("successfully retrieved user transactions detail", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, resp)
}
