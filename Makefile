proto:
	protoc -I/Users/gajduler/homebrew/include -I./pb \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:./pb \
		pb/auth.proto
	protoc -I/Users/gajduler/homebrew/include -I./pb \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:./pb \
		pb/auth.proto
serve:
	go run main.go serve