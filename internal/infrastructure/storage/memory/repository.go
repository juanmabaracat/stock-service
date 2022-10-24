package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/juanmabaracat/stock-service/internal/domain/stockinformation"
)

type Repository struct {
	stock map[string]stockinformation.StockInformation
}

func NewRepository() Repository {
	return Repository{stock: make(map[string]stockinformation.StockInformation)}
}

func (r Repository) GetByCode(productCode uuid.UUID) (*stockinformation.StockInformation, error) {
	stock, ok := r.stock[productCode.String()]
	if !ok {
		return nil, nil
	}

	return &stock, nil
}

func (r Repository) Add(stock stockinformation.StockInformation) error {
	r.stock[stock.ProductCode.String()] = stock
	return nil
}

func (r Repository) Update(stock stockinformation.StockInformation) error {
	r.stock[stock.ProductCode.String()] = stock
	return nil
}

func (r Repository) Delete(productCode uuid.UUID) error {
	_, found := r.stock[productCode.String()]
	if !found {
		return fmt.Errorf("product code %s not found", productCode.String())
	}

	delete(r.stock, productCode.String())
	return nil
}
