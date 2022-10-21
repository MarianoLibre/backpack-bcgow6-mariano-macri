package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock       []Product
	readWasCalled  bool
	writeWasCalled bool
	errWrite       string
	errRead        string
}

func (m *MockStorage) Read(data interface{}) error {
	m.readWasCalled = true
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]Product)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	m.writeWasCalled = true
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]Product)
	m.dataMock = a
	return nil
}

// Ejercicio 1 - Service/Repo/Db Update()
func TestServiceUpdate(t *testing.T) {
	product := Product{
		Id:        1,
		Name:      "tablet",
		Colour:    "black",
		Price:     200.5,
		Stock:     12,
		Code:      "XXX",
		Published: true,
		CreatedAt: "today",
	}
	updated := Product{
		Id:        1,
		Name:      "cellphone",
		Colour:    "red",
		Price:     500.5,
		Stock:     10,
		Code:      "XXX",
		Published: true,
		CreatedAt: "today",
	}
	database := []Product{
		product,
	}
	mockedStorage := MockStorage{
		dataMock: database,
	}
	rp := NewRepository(&mockedStorage)
	svc := NewService(rp)

	result, err := svc.Update(
		updated.Id, 
		updated.Name, 
		updated.Colour, 
		updated.Code, 
		updated.CreatedAt, 
		updated.Stock, 
		updated.Price, 
		updated.Published)

	assert.Nil(t, err)
	assert.Equal(t, result, updated)
	assert.True(t, mockedStorage.readWasCalled)
}

// Ejercicio 2 - Service/Repo/Db Delete()
func TestServiceDelete(t *testing.T) {
	p1 := Product{
			Id: 1,
			Name: "tablse",
			Colour: "black",
			Price: 200.5,
			Stock: 12,
			Code: "XXX",
			Published: true,
			CreatedAt: "today",
		}
	p2 := Product{
			Id: 2,
			Name: "celu",
			Colour: "red",
			Price: 500.5,
			Stock: 12,
			Code: "ZZZ",
			Published: true,
			CreatedAt: "today",
		}
	database := []Product{p1, p2}

	mockedStorage := MockStorage{dataMock: database}
	rp := NewRepository(&mockedStorage)
	svc := NewService(rp)

	err := svc.Delete(2)
	assert.Nil(t, err)
	assert.True(t, mockedStorage.readWasCalled)
	assert.True(t, mockedStorage.writeWasCalled)
	assert.Equal(t, mockedStorage.dataMock, []Product{p1})

	err = svc.Delete(2)
	assert.NotNil(t, err)
	assert.True(t, mockedStorage.readWasCalled)
	assert.True(t, mockedStorage.writeWasCalled)
	assert.Equal(t, mockedStorage.dataMock, []Product{p1})
}
