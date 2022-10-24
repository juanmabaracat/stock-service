package stockinformation

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (repo MockRepository) GetByCode(productCode uuid.UUID) (*StockInformation, error) {
	args := repo.Called(productCode)
	return args.Get(0).(*StockInformation), args.Error(1)
}

func (repo MockRepository) Add(stock StockInformation) error {
	args := repo.Called(stock)
	return args.Error(0)
}

func (repo MockRepository) Update(stock StockInformation) error {
	args := repo.Called(stock)
	return args.Error(0)
}

func (repo MockRepository) Delete(productCode uuid.UUID) error {
	args := repo.Called(productCode)
	return args.Error(0)
}
