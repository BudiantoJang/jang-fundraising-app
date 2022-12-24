package delivery

import (
	"fmt"
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

func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		resp := helper.APIResponse("failed uploading campaign image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := helper.APIResponse("failed uploading campaign image", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User.ID = currentUser.ID
	path := fmt.Sprintf("images/%d-%s", currentUser.ID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := helper.APIResponse("failed uploading campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	_, err = h.usecase.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := helper.APIResponse("failed saving campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data := gin.H{"is_uploaded": true}
	resp := helper.APIResponse("campaign image successfully updated", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, resp)
}
