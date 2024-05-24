package http

import (
	"errors"
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"fase-4-hf-product/mocks"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_Handler$
func Test_Handler(t *testing.T) {

	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name            string
		args            args
		mockApplication mocks.MockApplication
		wantOut         string
		isWantedErr     bool
	}{
		{
			name: "success_getByCategory",
			args: args{
				method: "GET",
				url:    "hermes_foods/product",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOutList: []dto.OutputProduct{
					{},
				},
				WantOut:     nil,
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     "[{}]",
			isWantedErr: false,
		},
		{
			name: "error_getByUUID",
			args: args{
				method: "GET",
				url:    "hermes_foods/product/100000",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut:     nil,
				WantErr:     errors.New("errGetProductByUUID"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "route GET hermes_foods/product/100000 not found"}`,
			isWantedErr: false,
		},
		{
			name: "success_save",
			args: args{
				method: "POST",
				url:    "hermes_foods/product",
				body:   strings.NewReader(`{"name":"Marty", "uuid":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"uuid":"0001","name":"Marty"}`,
			isWantedErr: false,
		},
		{
			name: "error_save_unmarshal",
			args: args{
				method: "POST",
				url:    "hermes_foods/product/",
				body:   strings.NewReader(`<=>`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to Unmarshal: invalid character '<' looking for beginning of value"}`,
			isWantedErr: true,
		},
		{
			name: "error_save",
			args: args{
				method: "POST",
				url:    "hermes_foods/product",
				body:   strings.NewReader(`{"name":"Marty", "uuid":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     errors.New("errSaveProduct"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to save product: errSaveProduct"}`,
			isWantedErr: false,
		},
		{
			name: "error_route_not_found",
			args: args{
				method: "PATCH",
				url:    "/hermes_foods/product",
				body:   strings.NewReader(`{"name":"Marty", "uuid":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     errors.New("errSaveProduct"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "route PATCH /hermes_foods/product not found"}`,
			isWantedErr: false,
		},
		{
			name: "success_update",
			args: args{
				method: "PUT",
				url:    "hermes_foods/product/1000",
				body:   strings.NewReader(`{"name":"Marty", "uuid":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"uuid":"0001","name":"Marty"}`,
			isWantedErr: false,
		},
		{
			name: "error_save_unmarshal",
			args: args{
				method: "PUT",
				url:    "hermes_foods/product/1000",
				body:   strings.NewReader(`<=>`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to Unmarshal: invalid character '<' looking for beginning of value"}`,
			isWantedErr: true,
		},
		{
			name: "success_delete",
			args: args{
				method: "DELETE",
				url:    "hermes_foods/product/1000",
				body:   strings.NewReader(`{"name":"Marty", "uuid":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"status":"OK"}`,
			isWantedErr: false,
		},
		{
			name: "error_delete",
			args: args{
				method: "DELETE",
				url:    "hermes_foods/product/",
				body:   strings.NewReader(`{"name":"Marty", "uuid":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputProduct{
					UUID:      "0001",
					Name:      "Marty",
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "route DELETE hermes_foods/product/ not found"}`,
			isWantedErr: false,
		},
	}

	for _, tc := range tests {
		h := NewHandler(tc.mockApplication)
		t.Run(tc.name, func(*testing.T) {

			req := httptest.NewRequest(tc.args.method, "/", tc.args.body)
			req.URL.Path = tc.args.url
			rw := httptest.NewRecorder()

			h.Handler(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}

		})
	}
}

// go test -v -count=1 -failfast -run ^Test_HealthCheck$
func Test_HealthCheck(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name            string
		args            args
		wantOut         string
		mockApplication mocks.MockApplication
		isWantedErr     bool
	}{
		{
			name: "success",
			args: args{
				method: "GET",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"status": "OK"}`,
			isWantedErr: false,
		},
		{
			name: "error_method_not_allowed",
			args: args{
				method: "POST",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"error": "method not allowed"}`,
			isWantedErr: true,
		},
	}

	for _, tc := range tests {
		h := NewHandler(tc.mockApplication)
		t.Run(tc.name, func(*testing.T) {
			req := httptest.NewRequest(tc.args.method, tc.args.url, tc.args.body)
			rw := httptest.NewRecorder()

			h.HealthCheck(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}
		})
	}
}
