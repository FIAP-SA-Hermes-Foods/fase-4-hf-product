package bdd

import (
	"context"
	"fase-4-hf-product/external/db/dynamo"
	l "fase-4-hf-product/external/logger"
	ps "fase-4-hf-product/external/strings"
	repositories "fase-4-hf-product/internal/adapters/driven/repositories/nosql"
	"fase-4-hf-product/internal/core/application"
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"fase-4-hf-product/internal/core/useCase"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/marcos-dev88/genv"
)

// go test -v -count=1 -failfast -run ^Test_GetProductByUUID$
func Test_GetProductByUUID(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST GetProductByUUID <====")

	type Input struct {
		ID string `json:"id"`
	}

	type Output struct {
		Output *dto.OutputProduct `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid and existing ID",
			name:     "success_valid_id",
			input: Input{
				ID: "b6b9a3a2-743b-45ff-92e9-df303b2af6a2",
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputProduct{
					Name:          "x-bacon",
					Category:      "meal",
					Image:         "url-image",
					Description:   "x-burguer with juicy meat, bread and cheese and so much bacon! yummy!",
					Price:         19.9,
					CreatedAt:     "19-05-2024 23:13:29",
					DeactivatedAt: "01-12-2024 15:58:21",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewProductRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewProductUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			product, err := app.GetProductByUUID(tc.input.ID)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if product.Name != tc.expectedOutput.Output.Name {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Name, product.Name)
				}

				if product.Category != tc.expectedOutput.Output.Category {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Category, product.Category)
				}

				if product.Image != tc.expectedOutput.Output.Image {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Image, product.Image)
				}

				if product.Description != tc.expectedOutput.Output.Description {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Description, product.Description)
				}

				if product.Price != tc.expectedOutput.Output.Price {
					t.Errorf("expected: %v\ngot: %v", tc.expectedOutput.Output.Price, product.Price)
				}

				if product.DeactivatedAt != tc.expectedOutput.Output.DeactivatedAt {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.DeactivatedAt, product.DeactivatedAt)
				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test GetProductByUUID <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_GetProductByCategory$
func Test_GetProductByCategory(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST GetProductByCategory <====")

	type Input struct {
		Category string `json:"category"`
	}

	type Output struct {
		Output []dto.OutputProduct `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid and existing category",
			name:     "success_valid_id",
			input: Input{
				Category: "meal",
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: []dto.OutputProduct{
					{
						Name:          "x-bacon",
						Category:      "meal",
						Image:         "url-image",
						Description:   "x-burguer with juicy meat, bread and cheese and so much bacon! yummy!",
						Price:         19.9,
						CreatedAt:     "19-05-2024 23:13:29",
						DeactivatedAt: "21-05-2024 21:03:38",
					},
					{
						Name:          "x-bacon",
						Category:      "meal",
						Image:         "url-image",
						Description:   "x-burguer with juicy meat, bread and cheese and so much bacon! yummy!",
						Price:         19.9,
						CreatedAt:     "19-05-2024 23:13:29",
						DeactivatedAt: "21-05-2024 19:46:06",
					},
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewProductRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewProductUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			productList, err := app.GetProductByCategory(tc.input.Category)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				for i := range productList {
					if productList[i].Name != tc.expectedOutput.Output[i].Name {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].Name, productList[i].Name)
					}

					if productList[i].Category != tc.expectedOutput.Output[i].Category {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].Category, productList[i].Category)
					}

					if productList[i].Image != tc.expectedOutput.Output[i].Image {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].Image, productList[i].Image)
					}

					if productList[i].Description != tc.expectedOutput.Output[i].Description {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].Description, productList[i].Description)
					}

					if productList[i].Price != tc.expectedOutput.Output[i].Price {
						t.Errorf("expected: %v\ngot: %v", tc.expectedOutput.Output[i].Price, productList[i].Price)
					}

					if productList[i].DeactivatedAt != tc.expectedOutput.Output[i].DeactivatedAt {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].DeactivatedAt, productList[i].DeactivatedAt)
					}
				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test GetProductByCategory <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_SaveProduct$
func Test_SaveProduct(t *testing.T) {
	t.Skip("skipping this test, comment it to test it properly")

	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST SaveProduct <====")

	type Input struct {
		Input *dto.RequestProduct `json:"input"`
	}

	type Output struct {
		Output *dto.OutputProduct `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid input",
			name:     "success",
			input: Input{
				Input: &dto.RequestProduct{
					Name:          "x-bacon",
					Category:      "meal",
					Image:         "url-image",
					Description:   "x-burguer with juicy meat, bread and cheese and so much bacon! yummy!",
					Price:         19.9,
					DeactivatedAt: "01-12-2024 15:58:21",
				},
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputProduct{
					Name:          "x-bacon",
					Category:      "meal",
					Image:         "url-image",
					Description:   "x-burguer with juicy meat, bread and cheese and so much bacon! yummy!",
					Price:         19.9,
					DeactivatedAt: "01-12-2024 15:58:21",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewProductRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewProductUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			product, err := app.SaveProduct(*tc.input.Input)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if product.Name != tc.expectedOutput.Output.Name {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Name, product.Name)
				}

				if product.Category != tc.expectedOutput.Output.Category {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Category, product.Category)
				}

				if product.Image != tc.expectedOutput.Output.Image {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Image, product.Image)
				}

				if product.Description != tc.expectedOutput.Output.Description {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Description, product.Description)
				}

				if product.Price != tc.expectedOutput.Output.Price {
					t.Errorf("expected: %v\ngot: %v", tc.expectedOutput.Output.Price, product.Price)
				}

				if product.DeactivatedAt != tc.expectedOutput.Output.DeactivatedAt {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.DeactivatedAt, product.DeactivatedAt)
				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test SaveProduct <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_UpdateProductByUUID$
func Test_UpdateProductByUUID(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST UpdateProductByUUID <====")

	type Input struct {
		UUID  string              `json:"uuid"`
		Input *dto.RequestProduct `json:"input"`
	}

	type Output struct {
		Output *dto.OutputProduct `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid input",
			name:     "success",
			input: Input{
				UUID: "ca32ec6d-7868-45ed-a905-01c4d0b2e62b",
				Input: &dto.RequestProduct{
					Name:          "x-burger",
					Category:      "meal",
					Description:   "x-burguer with juicy meat, bread and cheese",
					CreatedAt:     "01-05-2024 15:58:21",
					Price:         19.9,
					DeactivatedAt: "01-12-2024 15:58:21",
				},
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputProduct{
					UUID:          "ea985699-f14a-42b7-be9c-1d2a63080a6b",
					Name:          "x-burger",
					Category:      "meal",
					Image:         "url-image",
					Description:   "x-burguer with juicy meat, bread and cheese",
					Price:         19.9,
					CreatedAt:     "01-05-2024 15:58:21",
					DeactivatedAt: "01-12-2024 15:58:21",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewProductRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewProductUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			product, err := app.UpdateProductByUUID(tc.input.UUID, *tc.input.Input)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if product.Name != tc.expectedOutput.Output.Name {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Name, product.Name)
				}

				if product.Category != tc.expectedOutput.Output.Category {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Category, product.Category)
				}

				if product.Image != tc.expectedOutput.Output.Image {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Image, product.Image)
				}

				if product.Description != tc.expectedOutput.Output.Description {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Description, product.Description)
				}

				if product.Price != tc.expectedOutput.Output.Price {
					t.Errorf("expected: %v\ngot: %v", tc.expectedOutput.Output.Price, product.Price)
				}

				if product.DeactivatedAt != tc.expectedOutput.Output.DeactivatedAt {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.DeactivatedAt, product.DeactivatedAt)
				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test UpdateProductByUUID <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_DeleteProductByUUID$
func Test_DeleteProductByUUID(t *testing.T) {
	t.Skip("skipping this test, comment it to test it properly")

	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST Test_DeleteProductByUUID <====")

	type Input struct {
		ID string `json:"id"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
	}{
		{
			scenario: "Sending a valid and existing ID",
			name:     "success_valid_id",
			input: Input{
				ID: "ea985699-f14a-42b7-be9c-1d2a63080a6b",
			},
			shouldReturnError: false,
			shouldBeNull:      false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewProductRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewProductUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", "null"))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			errDelete := app.DeleteProductByUUID(tc.input.ID)

			if (!tc.shouldReturnError) && errDelete != nil {
				t.Errorf("unexpected error: %v", errDelete)
			}

			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test Test_DeleteProductByUUID <====")
	}
}
