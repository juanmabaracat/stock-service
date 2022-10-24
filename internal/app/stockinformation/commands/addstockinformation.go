package commands

import (
	"github.com/google/uuid"
	"github.com/juanmabaracat/stock-service/internal/domain/stockinformation"
	id "github.com/juanmabaracat/stock-service/internal/pkg/uuid"
)

type AddStockRequest struct {
	ProductName   string
	StockQuantity int64
}

type AddStockHandler interface {
	Handle(command AddStockRequest) (*uuid.UUID, error)
}

type addStockHandler struct {
	uuidProvider id.Provider
	repository   stockinformation.Repository
}

func NewAddStockHandler(uuidProvider id.Provider, repo stockinformation.Repository) AddStockHandler {
	return addStockHandler{
		uuidProvider: uuidProvider,
		repository:   repo,
	}
}

func (h addStockHandler) Handle(request AddStockRequest) (*uuid.UUID, error) {
	newID := h.uuidProvider.NewUUID()
	stockInfo := stockinformation.StockInformation{
		ProductCode:   newID,
		ProductName:   request.ProductName,
		StockQuantity: request.StockQuantity,
	}

	if err := h.repository.Add(stockInfo); err != nil {
		return nil, err
	}

	return &newID, nil
}
