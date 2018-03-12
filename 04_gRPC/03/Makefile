all: generate clean output

generate:
	protoc -I api/ api/api.proto --go_out=plugins=grpc:api

output:
	cp -r * $(GOPATH)/src/github.com/lukaskellerstein/Microservices/04_gRPC/03/

clean:
	rm -rf $(GOPATH)/src/github.com/lukaskellerstein/Microservices/04_gRPC/03/*