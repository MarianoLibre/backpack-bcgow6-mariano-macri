package loan

import "library.com/internal/domain"

type Service interface {
	GetAll() ([]domain.Loan, error)
	Get(int) (domain.Loan, error)
	Save(domain.Loan) (int, error)
	Update(domain.Loan) error
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

func (s *service) GetAll() ([]domain.Loan, error) {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (domain.Loan, error) {
	return s.repository.Get(id)
}

func (s *service) Save(loan domain.Loan) (int, error) {
	return s.repository.Save(loan)
}

func (s *service) Update(loan domain.Loan) error {
	return s.repository.Update(loan)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
