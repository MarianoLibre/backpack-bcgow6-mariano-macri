package product

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("product not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
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

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	productList, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Exists(ctx context.Context, productCode string) bool {
	exist := s.repository.Exists(ctx, productCode)
	return exist
}

func (s *service) Save(ctx context.Context, product domain.Product) (int, error) {
	lastId, err := s.repository.Save(ctx, product)
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *service) Update(ctx context.Context, p domain.Product) error {
	err := s.repository.Update(ctx, p)
	if err != nil {
		return err
	}
	return err
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
