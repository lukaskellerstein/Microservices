

Create traefik network 

`docker network create --driver overlay traefik-net`


Build docker images for apps

```Shell
docker build -t myapp1 ./app1
docker build -t myapp2 ./app2
docker build -t myapp3 ./app3
```
 
Run docker stack

```Shell
docker stack deploy -c docker-stack.yml testtraefikstack
```

Edit /etc/hosts - Local DNS
add local IP address and all enpoints from docker-stack.yml

```Shell
127.0.0.1	localhost
127.0.1.1	cellarstone-XPS-15-9560
192.168.1.19   app1.traefik
192.168.1.19   app2.traefik
192.168.1.19   app3.traefik

# The following lines are desirable for IPv6 capable hosts
::1     ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
```

Open apps

App3 
`http://app3.traefik/`

App2 
`http://app2.traefik/hamburger`

App1 
`http://app1.traefik/all`



