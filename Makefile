proto:
	protoc pkg/auth/pb/*.proto --go_out=plugins=grpc:.

client:
	go run cmd/main.go