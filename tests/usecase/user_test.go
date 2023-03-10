package tests

import (
	"jangFundraising/user"
	"log"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_User_RegisterUser(t *testing.T) {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	repo := user.NewRepository(db)
	useCase := user.NewUsecase(repo)

	t.Run("inserting new user must not return error", func(t *testing.T) {
		newUser := user.RegisterUserInput{
			Name:       faker.New().Person().Name(),
			Occupation: faker.New().Company().JobTitle(),
			Email:      faker.New().Internet().Email(),
			Password:   faker.New().Internet().Password(),
		}
		_, err := useCase.RegisterUser(newUser)
		assert.Nil(t, err)
	})
}

func Test_User_VerifyLogin(t *testing.T) {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	repo := user.NewRepository(db)
	useCase := user.NewUsecase(repo)

	testUser, _ := repo.FindByID(1)

	t.Run("invalid user credential should return error", func(t *testing.T) {

		loginData := user.LoginUserInput{
			Email:    testUser.Email,
			Password: faker.New().Internet().Password(),
		}

		_, err = useCase.VerifyLogin(loginData)
		assert.Error(t, err)
	})

	t.Run("valid user credential should not return error", func(t *testing.T) {
		newUser := user.RegisterUserInput{
			Name:       faker.New().Person().Name(),
			Occupation: faker.New().Company().JobTitle(),
			Email:      faker.New().Internet().Email(),
			Password:   faker.New().Internet().Password(),
		}
		_, err := useCase.RegisterUser(newUser)
		assert.Nil(t, err)

		loginData := user.LoginUserInput{
			Email:    newUser.Email,
			Password: newUser.Password,
		}

		_, err = useCase.VerifyLogin(loginData)
		assert.Nil(t, err)
	})
}

func Test_User_IsEmailAvailable(t *testing.T) {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	repo := user.NewRepository(db)
	useCase := user.NewUsecase(repo)

	testUser, _ := repo.FindByID(1)

	emailCheck := user.CheckEmailInput{
		Email: testUser.Email,
	}

	_, err = useCase.IsEmailAvailable(emailCheck)
	assert.Nil(t, err)
}
