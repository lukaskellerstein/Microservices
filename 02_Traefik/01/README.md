# Local Docker variant

Build docker images for apps

```Shell
docker build -t myapp1 ./app1
docker build -t myapp2 ./app2
```
 
Run docker compose

```Shell
docker-compose build
docker-compose up
```

Open apps

Traefic dashboard
`http://localhost:8080`

Whoami - (works only in chrome)
`http://whoami.docker.localhost/`

App2 - (works only in chrome)
`http://app2.docker.localhost/hamburger`

App1 - (works everywhere)
`http://localhost/app1/all`
`http://<localIP>/app1/all`
`http://docker.localhost/app1/all`



