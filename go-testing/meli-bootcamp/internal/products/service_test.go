package products

import "fmt"

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
	m.dataMock = append(m.dataMock, a[len(a)-1])
	return nil
}
