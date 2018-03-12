

```Shell
docker build -t myrabbitmq ./rabbitmq

docker run -d -p 15672:15672 -p 5672:5672 -t myrabbitmq
```

