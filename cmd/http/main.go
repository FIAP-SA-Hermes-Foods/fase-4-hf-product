package main

import (
	"context"
	"fase-4-hf-product/external/db/dynamo"
	l "fase-4-hf-product/external/logger"
	repositories "fase-4-hf-product/internal/adapters/driven/repositories/nosql"
	"fase-4-hf-product/internal/core/application"
	"fase-4-hf-product/internal/core/useCase"
	httpH "fase-4-hf-product/internal/handler/http"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	db := dynamo.NewDynamoDB(cfg)

	repo := repositories.NewProductRepository(db, os.Getenv("DB_TABLE"))

	uc := useCase.NewProductUseCase()

	app := application.NewApplication(repo, uc)

	h := httpH.NewHandler(app)

	router.Handle("/hermes_foods/health", http.StripPrefix("/", httpH.Middleware(h.HealthCheck)))
	router.Handle("/hermes_foods/product/", http.StripPrefix("/", httpH.Middleware(h.Handler)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
