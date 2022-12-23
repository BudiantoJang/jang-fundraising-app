package delivery

import (
	"jangFundraising/campaign"
	"jangFundraising/helper"
	"jangFundraising/user"
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

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}
		resp := helper.APIResponse("failed creating new campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.usecase.CreateCampaign(input)
	if err != nil {
		resp := helper.APIResponse("failed creating new campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("success creating new campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, resp)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		resp := helper.APIResponse("failed updating campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var inputData campaign.CampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}
		resp := helper.APIResponse("failed updating campaign detail", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User.ID = currentUser.ID

	updatedCampaign, err := h.usecase.Update(inputID, inputData)
	if err != nil {
		resp := helper.APIResponse("failed updating campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("success updating campaign details", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, resp)
}
