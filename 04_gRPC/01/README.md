# Install Protobuf on Ubuntu

```Shell
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip

# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

# Optional: change owner
sudo chwon cellarstone /usr/local/bin/protoc
sudo chwon -R cellarstone /usr/local/include/google
```

# Install Protobuf Golang plugin 

`sudo apt-get install golang-goprotobuf-dev`


# Install gRPC

`go get google.golang.org/grpc`


# Generate API


Run for generate API in golang language from *.proto file

```Shell
protoc -I api/ \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:api \
    api/api.proto
```

`protoc -I api/ api/api.proto --go_out=plugins=grpc:api`

`protoc --go_out=plugins=grpc api/api.proto`

