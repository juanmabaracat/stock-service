package stockinformation

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/juanmabaracat/stock-service/internal/app"
	"github.com/juanmabaracat/stock-service/internal/app/stockinformation/commands"
	"github.com/juanmabaracat/stock-service/internal/app/stockinformation/queries"
	"log"
	"net/http"
)

const ProductCodeURLParam = "productCode"

type Handler struct {
	stockServices app.StockServices
}

func NewHandler(stockServices app.StockServices) *Handler {
	return &Handler{stockServices: stockServices}
}

type CreateStockRequest struct {
	ProductName   string `json:"product_name"`
	StockQuantity int64  `json:"stock_quantity"`
}

type CreateStockResponse struct {
	ProductCode string `json:"product_code"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var stockToAdd CreateStockRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&stockToAdd)

	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(decodeErr.Error())
		fmt.Fprintf(w, "wrong body")
		return
	}

	if stockToAdd.ProductName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "name can't be empty")
		return
	}

	if stockToAdd.StockQuantity < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "stock quantity can't be negative")
		return
	}

	id, handlerErr := h.stockServices.Commands.AddStockHandler.Handle(commands.AddStockRequest{
		ProductName:   stockToAdd.ProductName,
		StockQuantity: stockToAdd.StockQuantity,
	})

	if handlerErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, handlerErr.Error())
		return
	}

	log.Printf("stock created with id: %s", id.String())
	stockResponse := CreateStockResponse{ProductCode: id.String()}

	jsonResp, err := json.Marshal(stockResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
	return
}

func (h Handler) GetByProductCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCode := vars[ProductCodeURLParam]
	code, parseErr := uuid.Parse(productCode)

	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid product code")
		return
	}

	stock, err := h.stockServices.Queries.GetStockHandler.Handle(
		queries.GetStockRequest{
			ProductCode: code,
		},
	)

	if stock == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(stock)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, encodeErr.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
