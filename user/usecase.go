package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	RegisterUser(input RegisterUserInput) (User, error)
	VerifyLogin(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (s *usecase) RegisterUser(input RegisterUserInput) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}
	user := User{
		Name:         input.Name,
		Occupation:   input.Occupation,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Role:         "user",
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return User{}, err
	}

	//generate JWT token

	return newUser, nil
}

func (s *usecase) VerifyLogin(input LoginUserInput) (User, error) {
	usr, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return usr, err
	}

	if usr.ID == 0 {
		return usr, errors.New("user can't be found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(input.Password))
	if err != nil {
		return usr, err
	}

	return usr, nil
}

func (s *usecase) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	usr, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	if usr.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (s *usecase) SaveAvatar(ID int, fileLocation string) (User, error) {
	usr, err := s.repository.FindByID(ID)
	if err != nil {
		return usr, err
	}

	usr.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(usr)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, err
}
