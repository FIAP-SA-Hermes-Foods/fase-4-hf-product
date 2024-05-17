package useCase

import (
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"log"
	"testing"
)

// go test -v -failfast -run ^Test_GetProductByUUID$
func Test_GetProductByUUID(t *testing.T) {
	type args struct {
		uuid string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				uuid: "1000000000",
			},
			wantErr: false,
		},
		{
			name: "not_valid_uuid",
			args: args{
				uuid: "",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewProductUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.GetProductByUUID(tc.args.uuid)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_GetProductByCategory$
func Test_GetProductByCategory(t *testing.T) {
	type args struct {
		category string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				category: "meal",
			},
			wantErr: false,
		},
		{
			name: "not_valid_category_null",
			args: args{
				category: "",
			},
			wantErr: true,
		},
		{
			name: "not_valid_category",
			args: args{
				category: "<==>",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewProductUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.GetProductByCategory(tc.args.category)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_SaveProduct$
func Test_SaveProduct(t *testing.T) {

	type args struct {
		reqProduct dto.RequestProduct
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				reqProduct: dto.RequestProduct{
					UUID:      "",
					Name:      "",
					Category:  "meal",
					CreatedAt: "",
				},
			},
			wantErr: false,
		},
		{
			name: "not_valid_category",
			args: args{
				reqProduct: dto.RequestProduct{
					UUID:      "",
					Name:      "",
					Category:  "",
					CreatedAt: "",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewProductUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.SaveProduct(tc.args.reqProduct)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_UpdateProductByUUID$
func Test_UpdateProductByUUID(t *testing.T) {

	type args struct {
		uuid       string
		reqProduct dto.RequestProduct
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				uuid: "1",
				reqProduct: dto.RequestProduct{
					UUID:      "",
					Name:      "",
					Category:  "meal",
					CreatedAt: "",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid_uuid",
			args: args{
				uuid:       "",
				reqProduct: dto.RequestProduct{},
			},
			wantErr: true,
		},
		{
			name: "not_valid_category",
			args: args{
				uuid: "1",
				reqProduct: dto.RequestProduct{
					UUID:      "",
					Name:      "",
					Category:  "",
					CreatedAt: "",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewProductUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.UpdateProductByUUID(tc.args.uuid, tc.args.reqProduct)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_UpdateProductByUUID$
func Test_DeleteProductByUUID(t *testing.T) {
	type args struct {
		uuid string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				uuid: "1",
			},
			wantErr: false,
		},
		{
			name: "invalid_uuid",
			args: args{
				uuid: "",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewProductUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.DeleteProductByUUID(tc.args.uuid)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}

}
