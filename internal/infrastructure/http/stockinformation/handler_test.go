package stockinformation

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/juanmabaracat/stock-service/internal/app"
	"github.com/juanmabaracat/stock-service/internal/app/stockinformation/commands"
	"github.com/juanmabaracat/stock-service/internal/app/stockinformation/queries"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAddStockHandler struct {
	Handler func(command commands.AddStockRequest) (*uuid.UUID, error)
}

func (h MockAddStockHandler) Handle(command commands.AddStockRequest) (*uuid.UUID, error) {
	return h.Handler(command)
}

type MockGetStockHandler struct {
	Handler func(query queries.GetStockRequest) (*queries.GetStockResult, error)
}

func (h MockGetStockHandler) Handle(query queries.GetStockRequest) (*queries.GetStockResult, error) {
	return h.Handler(query)
}

func TestHandler_Create(t *testing.T) {
	mockUUID := uuid.MustParse("ea6c836c-52f9-11ed-bdc3-0242ac120002")

	tests := []struct {
		name         string
		command      app.Commands
		request      *http.Request
		result       string
		resultStatus int
	}{
		{
			name: "should create stock successfully",
			command: app.Commands{AddStockHandler: MockAddStockHandler{
				func(command commands.AddStockRequest) (*uuid.UUID, error) {
					return &mockUUID, nil
				}},
			},
			request:      createRequest("POST", CreateStockRequest{"Sprite 2L", 200}),
			result:       `{"product_code":"ea6c836c-52f9-11ed-bdc3-0242ac120002"}`,
			resultStatus: http.StatusCreated,
		},
		{
			name:         "should return bad request on invalid stock quantity",
			command:      app.Commands{},
			request:      createRequest("POST", CreateStockRequest{"Sprite 2L", -10}),
			result:       "stock quantity can't be negative",
			resultStatus: http.StatusBadRequest,
		},
		{
			name:         "should return bad request on empty product name",
			command:      app.Commands{},
			request:      createRequest("POST", CreateStockRequest{"", 100}),
			result:       "name can't be empty",
			resultStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			services := app.StockServices{Commands: tt.command}
			h := NewHandler(services)
			response := httptest.NewRecorder()
			h.Create(response, tt.request)
			assert.Contains(t, response.Body.String(), tt.result)
			assert.Equal(t, tt.resultStatus, response.Code)
		})
	}
}

func TestHandler_GetByProductCode(t *testing.T) {
	mockUUID := uuid.MustParse("ea6c836c-52f9-11ed-bdc3-0242ac120002")
	mockStock := &queries.GetStockResult{
		ProductCode:   mockUUID,
		ProductName:   "Coca-Cola",
		StockQuantity: 200,
	}

	tests := []struct {
		name         string
		query        app.Queries
		request      *http.Request
		result       string
		resultStatus int
	}{
		{
			name: "should return valid stock information",
			query: app.Queries{GetStockHandler: MockGetStockHandler{
				func(query queries.GetStockRequest) (*queries.GetStockResult, error) {
					return mockStock, nil
				}}},
			request:      createRequestWithParam("GET", nil, "ea6c836c-52f9-11ed-bdc3-0242ac120002"),
			result:       `{"ProductCode":"ea6c836c-52f9-11ed-bdc3-0242ac120002","ProductName":"Coca-Cola","StockQuantity":200}`,
			resultStatus: http.StatusOK,
		},
		{
			name: "should return not found",
			query: app.Queries{GetStockHandler: MockGetStockHandler{
				func(query queries.GetStockRequest) (*queries.GetStockResult, error) {
					return nil, nil
				}}},
			request:      createRequestWithParam("GET", nil, "ea6c836c-52f9-11ed-bdc3-0242ac120002"),
			result:       "Not Found",
			resultStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			services := app.StockServices{Queries: tt.query}
			h := NewHandler(services)
			response := httptest.NewRecorder()
			h.GetByProductCode(response, tt.request)
			assert.Contains(t, response.Body.String(), tt.result)
			assert.Equal(t, tt.resultStatus, response.Code)
		})
	}
}

func createRequest(method string, data interface{}) *http.Request {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(data)
	req := httptest.NewRequest(method, "/stock", body)
	return req
}

func createRequestWithParam(method string, data interface{}, param string) *http.Request {
	req := createRequest(method, data)
	req = mux.SetURLVars(req, map[string]string{ProductCodeURLParam: param})
	return req
}
