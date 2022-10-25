package commands

import (
	"errors"
	"github.com/google/uuid"
	"github.com/juanmabaracat/stock-service/internal/domain/stockinformation"
	id "github.com/juanmabaracat/stock-service/internal/pkg/uuid"
	"reflect"
	"testing"
)

func TestNewAddStockInformationHandler(t *testing.T) {
	t.Run("should create AddStockInformationHandler successfully", func(t *testing.T) {
		expected := addStockHandler{
			uuidProvider: id.MockProvider{},
			repository:   stockinformation.MockRepository{},
		}

		got := NewAddStockHandler(id.MockProvider{}, stockinformation.MockRepository{})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("NewAddStockHandler() = %v, want %v", got, expected)
		}
	})
}

func Test_addStockInformationHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("ea6c836c-52f9-11ed-bdc3-0242ac120002")

	stockInfo := stockinformation.StockInformation{
		ProductCode:   mockUUID,
		ProductName:   "Chocolate Bar 50g",
		StockQuantity: 100,
	}

	type fields struct {
		uuidProvider id.Provider
		repository   stockinformation.Repository
	}
	type args struct {
		request AddStockRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *uuid.UUID
		wantErr bool
	}{
		{
			name: "should add the stock info and return its id",
			fields: fields{
				uuidProvider: newMockProvider(mockUUID),
				repository:   newMockRepository(stockInfo, nil),
			},
			args: args{request: AddStockRequest{
				ProductName:   "Chocolate Bar 50g",
				StockQuantity: 100,
			}},
			want:    &mockUUID,
			wantErr: false,
		},
		{
			name: "should return an error when trying to add new stock info",
			fields: fields{
				uuidProvider: newMockProvider(mockUUID),
				repository:   newMockRepository(stockInfo, errors.New("error adding stock")),
			},
			args: args{request: AddStockRequest{
				ProductName:   "Chocolate Bar 50g",
				StockQuantity: 100,
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &addStockHandler{
				uuidProvider: tt.fields.uuidProvider,
				repository:   tt.fields.repository,
			}
			got, err := handler.Handle(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func newMockProvider(mockUUID uuid.UUID) id.MockProvider {
	mockProv := id.MockProvider{}
	mockProv.On("NewUUID").Return(mockUUID)
	return mockProv
}

func newMockRepository(stock stockinformation.StockInformation, returnArg error) stockinformation.MockRepository {
	repo := stockinformation.MockRepository{}
	repo.On("Add", stock).Return(returnArg)
	return repo
}
