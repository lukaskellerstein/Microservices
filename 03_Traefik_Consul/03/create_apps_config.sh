
# WHOAMI ------------------------------------------------------------

# backend-1
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend1/circuitbreaker/expression -d 'NetworkErrorRatio() > 0.5'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend1/servers/server1/url -d 'http://whoami1:80'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend1/servers/server1/weight -d '10'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend1/servers/server2/url -d 'http://whoami2:80'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend1/servers/server2/weight -d '1'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend1/servers/server2/tags -d 'api,helloworld'

# backend-2
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/maxconn/amount -d '10'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/maxconn/extractorfunc -d 'request.host'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/loadbalancer/method -d 'drr'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/servers/server1/url -d 'http://whoami3:80'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/servers/server1/weight -d '1'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/servers/server2/url -d 'http://whoami4:80'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/servers/server2/weight -d '2'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend2/servers/server2/tags -d 'web'

# frontend-1
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend1/backend -d 'backend2'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend1/routes/test_1/rule -d 'Host:test.localhost'

# frontend-2
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend2/backend -d 'backend1'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend2/passHostHeader -d 'true'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend2/priority -d '10'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend2/entrypoints -d 'http'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend2/routes/test_2/rule -d 'PathPrefix:/test'


# FLASK APP ------------------------------------------------------------

# backends
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/amount -d '5'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/maxconn/extractorfunc -d 'request.host'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/loadbalancer/method -d 'drr'

curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/servers/server1/url -d 'http://whoami5:5000'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/servers/server1/weight -d '1'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/servers/server1/tags -d 'flask'

curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/servers/server2/url -d 'http://whoami6:5000'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/servers/server2/weight -d '2'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/backends/backend3/servers/server2/tags -d 'flask'

# frontend:
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend3/backend -d 'backend3'
curl -XPUT http://127.0.0.1:8500/v1/kv/traefik/frontends/frontend3/routes/test_1/rule -d 'Host:flask.localhost'
