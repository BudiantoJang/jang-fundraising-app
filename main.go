package main

import (
	"jangFundraising/auth"
	"jangFundraising/campaign"
	"jangFundraising/delivery"
	"jangFundraising/helper"
	"jangFundraising/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	databaseConn, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer databaseConn.Close()

	//Middleware
	authUsecase := auth.NewUsecase()

	//User
	userRepository := user.NewRepository(db)
	userUsecase := user.NewUsecase(userRepository)
	userHandler := delivery.NewUserHandler(userUsecase, *authUsecase)

	// Campaign

	campaignRepository := campaign.NewRepository(db)
	campaignUsecase := campaign.NewUsecase(campaignRepository)
	campaignHandler := delivery.NewCampaignHandler(campaignUsecase)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	user := api.Group("/user")
	user.POST("/", userHandler.RegisterUser)
	user.POST("/sessions", userHandler.Login)
	user.POST("/email_check", userHandler.CheckEmailAvailability)
	user.POST("/avatar", authMiddleware(*authUsecase, userRepository), userHandler.UploadAvatar)

	campaign := api.Group("/campaign")
	campaign.GET("/", campaignHandler.GetCampaigns)
	campaign.GET("/:id", campaignHandler.GetCampaignDetail)
	campaign.POST("/", authMiddleware(*authUsecase, userRepository), campaignHandler.CreateCampaign)

	router.Run()
}

func authMiddleware(autService auth.JWTService, userService user.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		bearer := strings.Split(authHeader, " ")
		stringToken := bearer[1]

		token, err := autService.ValidateToken(stringToken)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(payload["user_id"].(float64))

		user, err := userService.FindByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
