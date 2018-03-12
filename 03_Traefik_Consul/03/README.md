# Docker Swarm

Inspired by : http://blog.ruanbekker.com/blog/2017/10/24/managing-traefik-configuration-with-consul-on-docker-swarm/


Create network 

`docker network create --driver=overlay appnet`


Deploy to Docker Swarm

```Shell
docker stack deploy --compose-file consul.yml kvstore
docker stack deploy --compose-file traefik.yml proxy
docker stack deploy --compose-file apps.yml apps
```

Save configuration into Consul

```Shell
sh create_traefik_config.sh
sh create_apps_config.sh
```



Open Consul
`http://127.0.0.1:8500/ui/`

Look into key-value dictionary


Open Apps

App1
`http://127.0.0.1/test`


Others . . . ???