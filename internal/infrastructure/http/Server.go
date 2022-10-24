package http

import (
	"github.com/gorilla/mux"
	"github.com/juanmabaracat/stock-service/internal/app"
	stock "github.com/juanmabaracat/stock-service/internal/infrastructure/http/stockinformation"
	"log"
	"net/http"
)

type Server struct {
	appServices app.Services
	router      *mux.Router
}

func NewServer(appServices app.Services) *Server {
	httpServer := &Server{appServices: appServices}
	httpServer.router = mux.NewRouter()
	httpServer.AddStockInformationHTTPRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

func (s *Server) AddStockInformationHTTPRoutes() {
	const stockHTTPRoutePath = "/stock"
	// Commands
	s.router.HandleFunc(stockHTTPRoutePath, stock.NewHandler(s.appServices.StockServices).Create).Methods("POST")

	// Queries
	s.router.HandleFunc(stockHTTPRoutePath+"/{"+stock.ProductCodeURLParam+"}",
		stock.NewHandler(s.appServices.StockServices).GetByProductCode).Methods("GET")
}

// ListenAndServe Starts listening for requests
func (s *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
