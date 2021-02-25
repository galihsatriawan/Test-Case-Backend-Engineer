package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository Repository
}

type Service interface {
	FindUserByID(ID int) (User, error)
	FindUserByUsername(email string) (User, error)
	RegisterUser(input RegisterInput) (User, error)
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) FindUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *service) FindUserByUsername(username string) (User, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	user := User{}
	user.Username = input.Username
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)
	user.NamaLengkap = input.NamaLengkap
	// Check availability username
	checkUsername, err := s.repository.FindByUsername(input.Username)

	if err != nil {
		return user, err
	}

	if checkUsername.ID != 0 {
		return user, errors.New("Username is not available")
	}

	newUser, err := s.repository.Create(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
