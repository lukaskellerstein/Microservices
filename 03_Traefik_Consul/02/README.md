# Docker Swarm

If you haven't already builded images for app1 and app2

```Shell
docker build -t myapp1 ./app1
docker build -t myapp2 ./app2
docker build -t reverse-proxy ./traefik
```




Create network 

`docker network create --driver=overlay sky-net`

Deploy to Docker Swarm 

`docker stack deploy -c docker-stack.yml test4343`


> WARNING - Use **127.0.0.1** instead of **localhost**

Open Consul
`http://127.0.0.1:8500/ui/`

And then you will se ALL services, because our apps are self-registered.

Open Traefik 
`http://127.0.0.1:8080/dashboard/`


Open Apps

App1
`http://127.0.0.1/app1/all`

App2
`http://127.0.0.1/app2/hamburger`