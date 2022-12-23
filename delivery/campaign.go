package delivery

import (
	"jangFundraising/campaign"
	"jangFundraising/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	usecase campaign.Usecase
}

func NewCampaignHandler(usecase campaign.Usecase) *campaignHandler {
	return &campaignHandler{usecase}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.usecase.GetCampaigns(userID)
	if err != nil {
		resp := helper.APIResponse("error getting campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	response := helper.APIResponse("list of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignDetail(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		resp := helper.APIResponse("failed getting campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	campaignDetail, err := h.usecase.GetCampaignByID(input)
	if err != nil {
		resp := helper.APIResponse("failed getting campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, resp)
}
