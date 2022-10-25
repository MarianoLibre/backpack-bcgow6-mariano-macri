package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const errMsg = "Oops!"

type mockedRepository struct {}

func NewMockedRepository() Repository {
	return &mockedRepository{}
}

func (mr *mockedRepository) GetAllBySeller(SellerID string) ([]Product, error) {
	return nil, errors.New(errMsg)
}

func Test_GetAllBySeller(t *testing.T) {
	// Repository already contains mocked data, no need to mock anything here...
	repo := NewRepository()
	svc := NewService(repo)

	// Neither Repository nor Service check for valid seller ID,
	// The only validation is done in the handler. 
	// so, "err" will always be nil... But, anyway (see below)
	data, err := svc.GetAllBySeller("FEX112AC")
	assert.Nil(t, err)
	assert.Contains(t, data, Product{
		ID: "mock",
		SellerID: "FEX112AC",
		Description: "generic product",
		Price: 123.55,
	})

	repo = NewMockedRepository()
	svc = NewService(repo)
	// This will return an err...
	data, err = svc.GetAllBySeller("WTF!")
	assert.Nil(t, data)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New(errMsg))
}
