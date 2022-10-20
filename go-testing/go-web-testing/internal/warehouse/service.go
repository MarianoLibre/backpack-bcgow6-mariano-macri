package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("warehouse not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Warehouse) error
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

func (s *service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) Exists(ctx context.Context, warehouseCode string) bool {
	return s.repository.Exists(ctx, warehouseCode)
}

func (s *service) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	return s.repository.Save(ctx, w)
}

func (s *service) Update(ctx context.Context, w domain.Warehouse) error {
	return s.repository.Update(ctx, w)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
