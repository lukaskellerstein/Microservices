FROM golang
COPY . $GOPATH/src/github.com/lukaskellerstein/apps
WORKDIR $GOPATH/src/github.com/lukaskellerstein/apps
RUN go-wrapper download
RUN go-wrapper install
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app2  

# helpers
RUN pwd
RUN ls -l


# -----------------------
# Stage 2
# -----------------------
FROM alpine:3.5
COPY --from=0 /go/src/github.com/lukaskellerstein/apps/app2 /app/app2
RUN chmod +x /app/app2

# helpers
RUN ls -l  

EXPOSE 10012
ENTRYPOINT /app/app2


