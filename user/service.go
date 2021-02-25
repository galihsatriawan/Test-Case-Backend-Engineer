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
	Login(input LoginInput) (User, error)
	Update(currentUser User, input UpdateInput) (User, error)
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
func (s *service) Update(currentUser User, input UpdateInput) (User, error) {
	if input.Password != "" {
		currentUser.Password = input.Password
	}
	if input.NamaLengkap != "" {
		currentUser.NamaLengkap = input.NamaLengkap
	}
	updatedUser, err := s.repository.Save(currentUser)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
func (s *service) Login(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
