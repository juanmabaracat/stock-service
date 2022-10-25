package queries

import (
	"github.com/google/uuid"
	"github.com/juanmabaracat/stock-service/internal/domain/stockinformation"
	"reflect"
	"testing"
)

func Test_getStockRequestHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("ea6c836c-52f9-11ed-bdc3-0242ac120002")

	stockInfo := &GetStockResult{
		ProductCode:   mockUUID,
		ProductName:   "Coca-Cola",
		StockQuantity: 700,
	}

	type fields struct {
		repository stockinformation.Repository
	}
	type args struct {
		query GetStockRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetStockResult
		wantErr bool
	}{
		{
			name: "should return a valid stock information",
			fields: fields{repository: func() stockinformation.MockRepository {
				repo := stockinformation.MockRepository{}
				repo.On("GetByCode", mockUUID).Return(&stockinformation.StockInformation{
					ProductCode:   mockUUID,
					ProductName:   "Coca-Cola",
					StockQuantity: 700,
				}, nil)
				return repo
			}()},
			args:    args{query: GetStockRequest{ProductCode: mockUUID}},
			want:    stockInfo,
			wantErr: false,
		},
		{
			name: "should return nil when stock is not found",
			fields: fields{repository: func() stockinformation.MockRepository {
				repo := stockinformation.MockRepository{}
				var stock *stockinformation.StockInformation
				repo.On("GetByCode", mockUUID).Return(stock, nil)
				return repo
			}()},
			args:    args{query: GetStockRequest{ProductCode: mockUUID}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getStockRequestHandler{
				repository: tt.fields.repository,
			}
			got, err := h.Handle(tt.args.query)
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
