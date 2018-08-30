FROM golang
ADD . $GOPATH/src/github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Api
WORKDIR $GOPATH/src/github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Api/server.d

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o myapi

# helpers
# RUN echo $GOPATH
RUN pwd
RUN ls -l


# -----------------------
# Stage 2
# -----------------------
FROM alpine:3.5
COPY --from=0 /go/src/github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Api/server.d/myapi /app/myapi
RUN chmod +x /app/myapi

# helpers
RUN ls -l  

ENTRYPOINT /app/myapi



