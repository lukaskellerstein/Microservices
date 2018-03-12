
# APP1 ------------------------------------------------------------

curl -XPUT http://127.0.0.1:8500/v1/kv/app1/port -d '10011'
curl -XPUT http://127.0.0.1:8500/v1/kv/app1/id -d 'myapp111'
curl -XPUT http://127.0.0.1:8500/v1/kv/app1/name -d 'myapp111'


# APP2 ------------------------------------------------------------

curl -XPUT http://127.0.0.1:8500/v1/kv/app2/port -d '10012'
curl -XPUT http://127.0.0.1:8500/v1/kv/app2/id -d 'myapp222'
curl -XPUT http://127.0.0.1:8500/v1/kv/app2/name -d 'myapp222'
