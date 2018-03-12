# 01



Create traefik network 

`docker network create traefik-net`

or 

`docker network create --driver overlay traefik-net`


or 

`docker network create --driver overlay --scope swarm traefik-net`


Deploy stack

```Shell
docker stack deploy -c docker-stack.yml testtraefic1
```

Open apps

