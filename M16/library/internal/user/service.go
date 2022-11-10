package user

import "library.com/internal/domain"


type Service interface {
	GetAll() ([]domain.User, error)
	Get(int) (domain.User, error)
	Save(domain.User) (int, error)
	Update(domain.User) error
	Delete(int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.User, error) {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (domain.User, error) {
	return s.repository.Get(id)
}

func (s *service) Save(user domain.User) (int, error) {
	return s.repository.Save(user)
}

func (s *service) Update(user domain.User) error {
	return s.repository.Update(user)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
