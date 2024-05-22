build-proto:
	protoc \
	--go_out=product_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=product_proto \
	--go-grpc_opt=paths=source_relative \
	product.proto

run-terraform:
	terraform -chdir=infrastructure/terraform init;
	terraform -chdir=infrastructure/terraform plan;
	terraform -chdir=infrastructure/terraform apply;

run-bdd:
	docker build -f ./infrastructure/docker/Dockerfile.go_app_bdd -t hf-product-bdd:latest .;
	docker run --rm --name hf-product-bdd hf-product-bdd:latest
	@docker rmi -f hf-product-bdd >/dev/null 2>&1
	@docker rm $$(docker ps -a -f status=exited -q) -f >/dev/null 2>&1
