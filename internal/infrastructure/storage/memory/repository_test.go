package memory

import (
	"github.com/google/uuid"
	"github.com/juanmabaracat/stock-service/internal/domain/stockinformation"
	"reflect"
	"testing"
)

func TestRepository_Add(t *testing.T) {
	mockUUID := uuid.MustParse("ea6c836c-52f9-11ed-bdc3-0242ac120002")
	type fields struct {
		stock map[string]stockinformation.StockInformation
	}
	type args struct {
		stock stockinformation.StockInformation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "should add stock information",
			fields: fields{stock: make(map[string]stockinformation.StockInformation)},
			args: args{stock: stockinformation.StockInformation{
				ProductCode:   mockUUID,
				ProductName:   "Coca Cola 1,5L",
				StockQuantity: 50,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				stock: tt.fields.stock,
			}
			if err := r.Add(tt.args.stock); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetByCode(t *testing.T) {
	mockUUID := uuid.MustParse("ea6c836c-52f9-11ed-bdc3-0242ac120002")

	stockInfo := stockinformation.StockInformation{
		ProductCode:   uuid.UUID{},
		ProductName:   "Coca Cola 1,5L",
		StockQuantity: 50,
	}

	type fields struct {
		stock map[string]stockinformation.StockInformation
	}
	type args struct {
		productCode uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *stockinformation.StockInformation
		wantErr bool
	}{
		{
			name: "should return the stock information",
			fields: fields{
				stock: func() map[string]stockinformation.StockInformation {
					stock := make(map[string]stockinformation.StockInformation)
					stock[mockUUID.String()] = stockInfo
					return stock
				}(),
			},
			args:    args{productCode: mockUUID},
			want:    &stockInfo,
			wantErr: false,
		},
		{
			name: "should return nil it doesn't exist",
			fields: fields{
				stock: make(map[string]stockinformation.StockInformation),
			},
			args:    args{productCode: mockUUID},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				stock: tt.fields.stock,
			}
			got, err := r.GetByCode(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
