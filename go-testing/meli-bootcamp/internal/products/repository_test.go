package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)



type StubStore struct{
	mockedData []Product
	readWasCalled bool
}

func (s *StubStore) Read(data interface{}) error {
	s.readWasCalled = true
	a, ok := data.(*[]Product)
	if !ok {
		return errors.New("it failed!")
	}
	*a = s.mockedData
	return nil
}

func (s *StubStore) Write(data interface{}) error{
	return nil
}


func TestGetAll(t *testing.T) {
	expected := []Product{
		{
			Id: 1,
			Name: "Before updated",
			Colour: "black",
			Price: 200.5,
			Stock: 12,
			Code: "XXX",
			Published: true,
			CreatedAt: "today",
		},
		{
			Id: 2,
			Name: "celu",
			Colour: "red",
			Price: 500.5,
			Stock: 12,
			Code: "ZZZ",
			Published: true,
			CreatedAt: "today",
		},
	}
	db := &StubStore{mockedData: expected}
	r := NewRepository(db)

	out, err := r.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expected, out)
}

func TestUpdateName(t *testing.T) {
	data := []Product{
		{
			Id: 1,
			Name: "Before updated",
			Colour: "black",
			Price: 200.5,
			Stock: 12,
			Code: "XXX",
			Published: true,
			CreatedAt: "today",
		},
	}
	expected := data[0]
	updatedName := "Updated Name"
	expected.Name = updatedName
	db := &StubStore{mockedData: data, readWasCalled: false}
	r := NewRepository(db)

	out, err := r.UpdateName(1, updatedName)

	assert.Nil(t, err)
	assert.Equal(t, expected, out)
	assert.True(t, db.readWasCalled)
}
