

```Shell
docker build -t myprometheus ./prometheus

docker run -d -p 9090:9090 -t myprometheus
```

