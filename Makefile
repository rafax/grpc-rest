proto:
	protoc -I ./pb ./pb/auth.proto --go_out=plugins=grpc:pb
client:
	go run main.go client
serve:
	go run main.go serve