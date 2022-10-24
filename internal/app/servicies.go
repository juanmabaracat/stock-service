package app

import (
	"github.com/juanmabaracat/stock-service/internal/app/stockinformation/commands"
	"github.com/juanmabaracat/stock-service/internal/app/stockinformation/queries"
	"github.com/juanmabaracat/stock-service/internal/domain/stockinformation"
	"github.com/juanmabaracat/stock-service/internal/pkg/uuid"
)

type Commands struct {
	AddStockHandler commands.AddStockHandler
}

type Queries struct {
	GetStockHandler queries.GetStockRequestHandler
}

type StockServices struct {
	Commands
	Queries
}

type Services struct {
	StockServices
}

func NewServices(repo stockinformation.Repository, up uuid.Provider) Services {
	return Services{
		StockServices: StockServices{
			Commands: Commands{
				AddStockHandler: commands.NewAddStockHandler(up, repo),
			},
			Queries: Queries{
				GetStockHandler: queries.NewGetStockRequestHandler(repo),
			},
		},
	}
}
