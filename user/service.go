package user

type Service interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	Store(user UserRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()

	return users, err
	// return s.repository.FindAll()
}

func (s *service) FindById(id int) (User, error) {
	user, err := s.repository.FindById(id)

	return user, err
}

func (s *service) Store(userRequest UserRequest) (User, error) {
	user := User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
	}
	newUser, err := s.repository.Store(user)

	return newUser, err
}
