package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllBySeller(t *testing.T) {
	// Repository already contains mocked data, no need to mock anything here...
	repo := NewRepository()
	svc := NewService(repo)

	// Neither Repository nor Service check for valid seller ID,
	// The only validation is done in the handler. 
	// so, "err" will always be nil...
	data, err := svc.GetAllBySeller("FEX112AC")
	assert.Nil(t, err)
	assert.Contains(t, data, Product{
		ID: "mock",
		SellerID: "FEX112AC",
		Description: "generic product",
		Price: 123.55,
	})
}
