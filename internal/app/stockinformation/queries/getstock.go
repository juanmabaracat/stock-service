package queries

import (
	"github.com/google/uuid"
	"github.com/juanmabaracat/stock-service/internal/domain/stockinformation"
)

type GetStockRequest struct {
	ProductCode uuid.UUID
}

type GetStockResult struct {
	ProductCode   uuid.UUID
	ProductName   string
	StockQuantity int64
}

type GetStockRequestHandler interface {
	Handle(query GetStockRequest) (*GetStockResult, error)
}

type getStockRequestHandler struct {
	repository stockinformation.Repository
}

func NewGetStockRequestHandler(repository stockinformation.Repository) GetStockRequestHandler {
	return getStockRequestHandler{
		repository: repository,
	}
}

func (h getStockRequestHandler) Handle(query GetStockRequest) (*GetStockResult, error) {
	stock, err := h.repository.GetByCode(query.ProductCode)

	if stock == nil && err == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	result := &GetStockResult{
		ProductCode:   stock.ProductCode,
		ProductName:   stock.ProductName,
		StockQuantity: stock.StockQuantity,
	}

	return result, nil
}
