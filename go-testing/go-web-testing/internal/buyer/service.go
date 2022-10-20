package buyer

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type Service interface {
	Get(ctx context.Context, id int) (domain.Buyer, error)
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Update(ctx context.Context, b domain.Buyer) (domain.Buyer, error) 
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, b domain.Buyer) (int, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(ctx context.Context, id int) (domain.Buyer, error) {
	buyer, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Buyer{}, err
	}

	return buyer, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	b, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return b, nil
}
func (s *service) Save(ctx context.Context, b domain.Buyer) (int, error) {
	id, err := s.repository.Save(ctx, b)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *service) Update(ctx context.Context, b domain.Buyer) (domain.Buyer, error) {
	buyer, err := s.Get(ctx, b.ID)
	if err != nil {
		return domain.Buyer{}, err
	}

	if b.CardNumberID == "" {
		b.CardNumberID = buyer.CardNumberID
	}
	if b.FirstName == ""{
		b.FirstName = buyer.FirstName
	}
	if b.LastName == "" {
		b.LastName = buyer.LastName
	}

	err = s.repository.Update(ctx, b)
	if err != nil {
		return domain.Buyer{}, err
	}
	return b, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Exists(ctx context.Context, cardNumberID string) bool {
	exist := s.repository.Exists(ctx, cardNumberID)
	return exist
}
