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

func Test_User_Save(t *testing.T) {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	t.Run("save user must not be error", func(t *testing.T) {
		newUser := user.User{
			Name:         faker.New().Person().Name(),
			Occupation:   faker.New().Company().JobTitle(),
			Email:        faker.New().Internet().Email(),
			PasswordHash: faker.New().Hash().SHA256(),
		}

		repo := user.NewRepository(db)
		_, err := repo.Save(newUser)
		assert.Nil(t, err)
	})
}

func Test_User_FindByEmail(t *testing.T) {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	t.Run("doesnt get the user by email, must return an error", func(t *testing.T) {
		repo := user.NewRepository(db)
		email := faker.New().Internet()
		_, err := repo.FindByEmail(email.Email())

		assert.Error(t, err)
	})

	t.Run("find by email must not be error", func(t *testing.T) {
		newUser := user.User{
			Name:         faker.New().Person().Name(),
			Occupation:   faker.New().Company().JobTitle(),
			Email:        faker.New().Internet().Email(),
			PasswordHash: faker.New().Hash().SHA256(),
		}

		repo := user.NewRepository(db)
		usr, _ := repo.Save(newUser)

		_, err = repo.FindByEmail(usr.Email)

		assert.Nil(t, err)
	})
}

func Test_User_FindByID(t *testing.T) {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	t.Run("doesnt get the user by id, must return an error", func(t *testing.T) {
		repo := user.NewRepository(db)
		_, err := repo.FindByID(faker.New().Int())

		assert.Error(t, err)
	})

	t.Run("find by id must not be error", func(t *testing.T) {
		newUser := user.User{
			Name:         faker.New().Person().Name(),
			Occupation:   faker.New().Company().JobTitle(),
			Email:        faker.New().Internet().Email(),
			PasswordHash: faker.New().Hash().SHA256(),
		}

		repo := user.NewRepository(db)
		usr, _ := repo.Save(newUser)

		_, err = repo.FindByID(usr.ID)

		assert.Nil(t, err)
	})
}

func Test_User_Update(t *testing.T) {
	dsn := "root:Pasuruan_123@tcp(127.0.0.1:3306)/fundraising?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	t.Run("update user must not returns error", func(t *testing.T) {
		newUser := user.User{
			Name:         faker.New().Person().Name(),
			Occupation:   faker.New().Company().JobTitle(),
			Email:        faker.New().Internet().Email(),
			PasswordHash: faker.New().Hash().SHA256(),
		}

		repo := user.NewRepository(db)
		_, err := repo.Save(newUser)
		assert.Nil(t, err)

		newUser.Name = faker.New().Person().Name()
		newUser.Occupation = faker.New().Company().JobTitle()

		_, err = repo.Update(newUser)

		assert.Nil(t, err)
	})
}
