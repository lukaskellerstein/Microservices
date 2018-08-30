FROM golang
ADD . $GOPATH/src/github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Iot
WORKDIR $GOPATH/src/github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Iot/server.d

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o myiot

# helpers
# RUN echo $GOPATH
RUN pwd
RUN ls -l


# -----------------------
# Stage 2
# -----------------------
FROM alpine:3.5
COPY --from=0 /go/src/github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Iot/server.d/myiot /app/myiot
RUN chmod +x /app/myiot

# helpers
RUN ls -l  

ENTRYPOINT /app/myiot



