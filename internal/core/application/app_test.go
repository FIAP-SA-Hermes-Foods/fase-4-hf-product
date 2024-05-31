package application

import (
	"errors"
	ps "fase-4-hf-product/external/strings"
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"fase-4-hf-product/mocks"
	"log"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_GetProductByID$
func Test_GetProductByID(t *testing.T) {
	type args struct {
		uuid string
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockProductRepository
		mockUseCase    mocks.MockProductUseCase
		wantOut        dto.OutputProduct
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				uuid: "10000000",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: &dto.ProductDB{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "success_null",
			args: args{
				uuid: "10000000",
			},
			mockRepository: mocks.MockProductRepository{
				WantOutNull: "nilGetProductByID",
				WantOut:     nil,
				WantErr:     nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "error_repository",
			args: args{
				uuid: "10000000",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: errors.New("errGetProductByID"),
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut:       dto.OutputProduct{},
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				uuid: "10000000",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: errors.New("errGetProductByID"),
			},
			wantOut:       dto.OutputProduct{},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {

			out, err := app.GetProductByID(tc.args.uuid)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetProductByCategory$
func Test_GetProductByCategory(t *testing.T) {
	type args struct {
		category string
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockProductRepository
		mockUseCase    mocks.MockProductUseCase
		wantOut        []dto.OutputProduct
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				category: "MEAL",
			},
			mockRepository: mocks.MockProductRepository{
				WantOutList: []dto.ProductDB{
					{
						UUID:          "001",
						Name:          "someUser",
						CreatedAt:     "2001-01-01 15:30:00",
						DeactivatedAt: "2001-01-01 15:30:00",
					},
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: []dto.OutputProduct{
				{
					UUID:          "001",
					Name:          "someUser",
					CreatedAt:     "2001-01-01 15:30:00",
					DeactivatedAt: "2001-01-01 15:30:00",
				},
			},

			isWantedError: false,
		},
		{
			name: "success_null",
			args: args{
				category: "MEAL",
			},
			mockRepository: mocks.MockProductRepository{
				WantOutNull: "nilGetProductByCategory",
				WantOut:     nil,
				WantErr:     nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut:       nil,
			isWantedError: false,
		},
		{
			name: "error_repository",
			args: args{
				category: "MEAL",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: errors.New("errGetProductByCategory"),
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut:       nil,
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				category: "MEAL",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: errors.New("errGetProductByCategory"),
			},
			wantOut:       nil,
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {

			out, err := app.GetProductByCategory(tc.args.category)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_SaveProduct$
func Test_SaveProduct(t *testing.T) {
	type args struct {
		req dto.RequestProduct
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockProductRepository
		mockUseCase    mocks.MockProductUseCase
		wantOut        dto.OutputProduct
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOutNull: "nilGetProductByID",
				WantOut: &dto.ProductDB{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "success_null",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOutNull: "nilSaveProduct",
				WantOut:     nil,
				WantErr:     nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},

		{
			name: "error_user_exists",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: &dto.ProductDB{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: errors.New("errSaveProduct"),
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_repository",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: errors.New("errSaveProduct"),
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_getUserByID",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: errors.New("errGetProductByID"),
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {
			out, err := app.SaveProduct(tc.args.req)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_UpdateProductByID$
func Test_UpdateProductByID(t *testing.T) {
	type args struct {
		uuid string
		req  dto.RequestProduct
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockProductRepository
		mockUseCase    mocks.MockProductUseCase
		wantOut        dto.OutputProduct
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				uuid: "00001",
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOutNull: "nilGetProductByID",
				WantOut: &dto.ProductDB{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "success_null",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "error_useCase",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: errors.New("errUpdateProductByID"),
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_repository",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: errors.New("errSaveProduct"),
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_UpdateProductByID",
			args: args{
				req: dto.RequestProduct{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: errors.New("errUpdateProductByID"),
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputProduct{
				UUID:      "001",
				Name:      "someUser",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {
			out, err := app.UpdateProductByID(tc.args.uuid, tc.args.req)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_DeleteProductByID$
func Test_DeleteProductByID(t *testing.T) {
	type args struct {
		uuid string
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockProductRepository
		mockUseCase    mocks.MockProductUseCase
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				uuid: "10000000",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: &dto.ProductDB{
					UUID:      "001",
					Name:      "someUser",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			isWantedError: false,
		},
		{
			name: "error_repository",
			args: args{
				uuid: "10000000",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: errors.New("errDeleteProductByID"),
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: nil,
			},
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				uuid: "10000000",
			},
			mockRepository: mocks.MockProductRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockProductUseCase{
				WantErr: errors.New("errDeleteProductByID"),
			},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {

			err := app.DeleteProductByID(tc.args.uuid)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}
