package main

import (
	"jangFundraising/auth"
	"jangFundraising/delivery"
	"jangFundraising/user"
	"log"

	"github.com/gin-gonic/gin"
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

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := delivery.NewUserHandler(userService, *authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()
}
