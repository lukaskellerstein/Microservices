FROM golang
COPY . $GOPATH/src/github.com/lukaskellerstein/apps
WORKDIR $GOPATH/src/github.com/lukaskellerstein/apps
RUN go-wrapper download
RUN go-wrapper install
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app1

# helpers
RUN pwd
RUN ls -l


# -----------------------
# Stage 2
# -----------------------
FROM alpine:3.5
MAINTAINER lukaskellerstein
COPY --from=0 /go/src/github.com/lukaskellerstein/apps/app1 /app/app1
RUN chmod +x /app/app1

# helpers
RUN ls -l  

# ENV PORT 10011
# EXPOSE 10011

ENTRYPOINT /app/app1


