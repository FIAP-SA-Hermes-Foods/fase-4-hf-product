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
