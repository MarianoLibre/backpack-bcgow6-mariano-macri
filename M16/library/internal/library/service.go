package library

import "library.com/internal/domain"

type Service interface {
	GetAll() ([]domain.Library, error)
	Get(int) (domain.Library, error)
	Save(domain.Library) (int, error)
	Update(domain.Library) error
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

func (s *service) GetAll() ([]domain.Library, error) {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (domain.Library, error) {
	return s.repository.Get(id)
}

func (s *service) Save(library domain.Library) (int, error) {
	return s.repository.Save(library)
}

func (s *service) Update(library domain.Library) error {
	return s.repository.Update(library)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
