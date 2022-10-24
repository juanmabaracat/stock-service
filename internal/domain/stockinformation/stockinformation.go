package stockinformation

import "github.com/google/uuid"

type StockInformation struct {
	ProductCode   uuid.UUID
	ProductName   string
	StockQuantity int64
}
