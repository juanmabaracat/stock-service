package stockinformation

import "github.com/google/uuid"

type Repository interface {
	GetByCode(productCode uuid.UUID) (*StockInformation, error)
	Add(stock StockInformation) error
	Update(stock StockInformation) error
	Delete(productCode uuid.UUID) error
}
