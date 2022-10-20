package employee

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("employee not found")
)

type Service interface {
	Delete(ctx context.Context, id int) (string, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Save(ctx context.Context, employee domain.Employee) (domain.Employee, error)
	Update(ctx context.Context, id int, employeeReceived domain.Employee) (domain.Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Delete(ctx context.Context, id int) (response string, err error) {
	if _, err := s.repository.Get(ctx, id); err != nil {
		return "ID not found", err
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return "delete failed", err
	}

	return "employee deleted", nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Employee, error) {
	employee, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	employees, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *service) Save(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	if s.repository.Exists(ctx, employee.CardNumberID) {
		return domain.Employee{
			ID: -1,
		}, errors.New("the employee card number ID must be unique")
	}

	employeeId, errSave := s.repository.Save(ctx, employee)
	if errSave != nil {
		return domain.Employee{
			ID: -2,
		}, errSave
	}

	employee, errGet := s.repository.Get(ctx, employeeId)
	if errGet != nil {
		return domain.Employee{
			ID: 0,
		}, ErrNotFound
	}

	return employee, nil
}

func (s *service) Update(ctx context.Context, id int, employeeReceived domain.Employee) (domain.Employee, error) {
	employee, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Employee{
			ID: 0,
		}, err
	}

	if s.repository.Exists(ctx, employeeReceived.CardNumberID) {
		return domain.Employee{
			ID: -2,
		}, errors.New("the employee card number ID must be unique")
	}

	switch {
	case employeeReceived.CardNumberID != "":
		employee.CardNumberID = employeeReceived.CardNumberID
	case employeeReceived.FirstName != "":
		employee.FirstName = employeeReceived.FirstName
	case employeeReceived.LastName != "":
		employee.LastName = employeeReceived.LastName
	case employeeReceived.WarehouseID != 0: // If user wants to desassociate employee from its warehouse, it can set this attribute to -1
		employee.WarehouseID = employeeReceived.WarehouseID
	}

	if s.repository.Update(ctx, employee) != nil {
		return domain.Employee{
			ID: -1,
		}, err
	}

	return employee, nil
}
