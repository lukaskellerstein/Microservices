
# Run Jaeger

```Shell
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
  -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
```

# Download dependencies

If not exists `Gopkg.toml` file run `dep init`.

```Shell
dep ensure
```

# Run Apps


## Formatter microservice

```Shell
go run *.go
```

## Publisher microservice

```Shell
go run *.go
```

## Clients app

```Shell
go run *.go SomeName
```


# See Jeager tracer

`http://localhost:16686`