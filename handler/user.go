package handler

import (
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
