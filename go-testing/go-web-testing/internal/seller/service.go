package seller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
	"github.com/gin-gonic/gin"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
)

type Service interface {
	Save(c gin.Context, seller domain.Seller) (domain.Seller, error)
	GetAll(c gin.Context) ([]domain.Seller, error)
	GetSellerById(c gin.Context) (domain.Seller, error)
	Delete(c gin.Context) error
	Update(c gin.Context, request domain.Seller) (domain.Seller, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(c gin.Context, seller domain.Seller) (domain.Seller, error) {
	if seller.CID == 0 {
		return domain.Seller{}, fmt.Errorf("To register a vendor the CId cannot be 0")
	}
	exist := s.repository.Exists(&c, seller.CID)
	if exist {
		return domain.Seller{}, fmt.Errorf("El Cid is already registered please try again")
	}
	id, err := s.repository.Save(&c, seller)
	if err != nil {
		return domain.Seller{}, err
	}
	seller.ID = id
	return seller, nil
}
func (s *service) GetAll(c gin.Context) ([]domain.Seller, error) {
	sellers, err := s.repository.GetAll(&c)
	if err != nil {
		return []domain.Seller{}, err
	}
	return sellers, nil
}
func (s *service) GetSellerById(c gin.Context) (domain.Seller, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return domain.Seller{}, err
	}
	seller, err := s.repository.Get(&c, id)
	if err != nil {
		return domain.Seller{}, err
	}
	return seller, nil
}
func (s *service) Delete(c gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	erro := s.repository.Delete(&c, id)
	if erro != nil {
		return erro
	}
	return nil
}
func (s *service) Update(c gin.Context, request domain.Seller) (domain.Seller, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return domain.Seller{}, err
	}
	seller, err := s.repository.Get(&c, id)
	if err != nil {
		return domain.Seller{}, err
	}
	if request.Address != "" {
		seller.Address = request.Address
	}
	if request.CompanyName != "" {
		seller.CompanyName = request.CompanyName
	}
	if request.Telephone != "" {
		seller.Telephone = request.Telephone
	}
	if request.CID != 0 {
		exist := s.repository.Exists(&c, request.CID)
		if exist {
			return domain.Seller{}, fmt.Errorf("This CId cannot be updated because it already exists in another provider")
		}
		seller.CID = request.CID
	}
	err = s.repository.Update(&c, seller)
	return seller, nil
}
