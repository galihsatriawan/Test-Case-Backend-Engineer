package user

type service struct {
	repository Repository
}

type Service interface {
	FindUserByID(ID int) (User, error)
	FindUserByUsername(email string) (User, error)
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
