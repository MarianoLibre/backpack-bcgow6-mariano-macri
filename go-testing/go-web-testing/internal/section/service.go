package section

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Save(ctx context.Context, sectionNumber, currentTemperature, minimunTemperature, currentCapacity, minimumCapacity, maximunCapacity, warehouseID, productTypeID int) (domain.Section, error)
	SectionNumberAlreadyExists(ctx context.Context, sectionNumber int) bool
	GetByID(ctx context.Context, id int) (domain.Section, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, section domain.Section) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Retrieve all records from db
func (s *service) GetAll(ctx context.Context) ([]domain.Section, error) {
	sections, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

// Save a new record in db
func (s *service) Save(ctx context.Context, sectionNumber, currentTemperature, minimunTemperature, currentCapacity, minimumCapacity, maximunCapacity, warehouseID, productTypeID int) (domain.Section, error) {

	section := domain.NewSection(sectionNumber, currentTemperature, minimunTemperature, currentCapacity, minimumCapacity, maximunCapacity, warehouseID, productTypeID)

	id, err := s.repository.Save(ctx, *section)
	if err != nil {
		return domain.Section{}, err
	}

	section.ID = id

	return *section, nil

}

// Verify if a section number already exists
func (s *service) SectionNumberAlreadyExists(ctx context.Context, sectionNumber int) bool {
	return s.repository.Exists(ctx, sectionNumber)
}

// Retrieve a section by ID
func (s *service) GetByID(ctx context.Context, id int) (domain.Section, error) {
	section, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Section{}, ErrNotFound
	}
	return section, nil

}

// Delete the record in db
func (s *service) Delete(ctx context.Context, id int) error {
	if err := s.repository.Delete(ctx, id); err != nil {
		return ErrNotFound
	}
	return nil
}

// Update the record in db
func (s *service) Update(ctx context.Context, section domain.Section) error {
	if err := s.repository.Update(ctx, section); err != nil {
		return err
	}
	return nil
}
