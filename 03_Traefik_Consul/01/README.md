# Docker Compose

Run docker compose

```Shell
docker-compose up
```

Open Consul
`http://localhost:8500/ui/`

And then you will se ALL services, because our apps are self-registered.

Open Traefik 
`http://localhost:8080/dashboard/`


Open Apps

App1
`http://localhost:10011/all`
`http://myapp1.docker.localhost:10011/all`
??? `http://myapp2.docker.localhost:10011/all` ??? WTF ???

App2
`http://localhost:10012/hamburger`
`http://myapp1.docker.localhost:10012/hamburger`
??? `http://myapp2.docker.localhost:10012/hamburger` ??? WTF ???