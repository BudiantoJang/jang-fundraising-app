package delivery

import (
	"fmt"
	"jangFundraising/helper"
	"jangFundraising/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errMessage := gin.H{"errors": errors}

		resp := helper.APIResponse("error binding user", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	usr, err := h.userService.RegisterUser(input)
	if err != nil {
		resp := helper.APIResponse("error while trying to register user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	formattedUser := user.FormatUser(usr)
	resp := helper.APIResponse("account has been created", http.StatusOK, "success", formattedUser)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errMessage := gin.H{"errors": errors}

		resp := helper.APIResponse("error binding user", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	verifiedUser, err := h.userService.VerifyLogin(input)
	if err != nil {
		resp := helper.APIResponse("error verifying user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	formattedUser := user.FormatUser(verifiedUser)
	resp := helper.APIResponse("login success", http.StatusOK, "success", formattedUser)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		resp := helper.APIResponse("failed binding the payload", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	isEmailAvailable, _ := h.userService.IsEmailAvailable(input)

	data := gin.H{
		"isAvailable": isEmailAvailable,
	}

	metaMessage := "email has already been registered"

	if isEmailAvailable {
		metaMessage = "email is available"
	}

	resp := helper.APIResponse(metaMessage, http.StatusOK, "succes", data)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := helper.APIResponse("Failed saving avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	//harusnya dari jwt
	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := helper.APIResponse("Failed saving avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := helper.APIResponse("Failed saving avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	data := gin.H{"is_uploaded": true}
	resp := helper.APIResponse("Avatar successfully updated", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, resp)
}
