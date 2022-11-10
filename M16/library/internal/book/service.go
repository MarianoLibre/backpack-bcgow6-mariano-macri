package book

import "library.com/internal/domain"


type Service interface {
	GetAll() ([]domain.Book, error)
	Get(int) (domain.Book, error)
	Save(domain.Book) (int, error)
	Update(domain.Book) error
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

func (s *service) GetAll() ([]domain.Book, error) {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (domain.Book, error) {
	return s.repository.Get(id)
}

func (s *service) Save(book domain.Book) (int, error) {
	return s.repository.Save(book)
}

func (s *service) Update(book domain.Book) error {
	return s.repository.Update(book)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
