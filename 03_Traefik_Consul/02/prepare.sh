docker build -t myapp1 ./app1
docker build -t myapp2 ./app2
docker build -t reverse-proxy ./traefik
docker network create --driver=overlay sky-net