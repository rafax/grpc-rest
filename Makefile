proto:
	protoc -I/usr/local/include -I./pb \
		-I${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:./pb \
		pb/auth.proto
    protoc -I/usr/local/include -I./pb \
		-I${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:./pb \
		pb/auth.proto
client:
	go run main.go client
serve:
	go run main.go serve