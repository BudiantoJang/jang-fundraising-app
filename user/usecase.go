package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	VerifyLogin(input LoginUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
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

func (s *service) VerifyLogin(input LoginUserInput) (User, error) {
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