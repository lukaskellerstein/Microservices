# Simple variant

Run docker compose

```Shell
docker-compose up
```

Open Consul
`http://localhost:8500/ui/`

And then you will se only Consul service. So let's add another with Postman

Send PUT requests to this endpoint `http://localhost:8500/v1/agent/service/register`

```JSON
{
	"ID": "myapp1",
	"Name": "myapp1",
	"Address": "myapp1",
	"Port": 10011,
	"check": {
		"id": "ping",
		"name": "HTTP api port on 10011",
		"http": "http://myapp1:10011/ping",
		"interval": "5s",
		"timeout": "1s"
	}
}
```


```JSON
{
	"ID": "myapp2",
	"Name": "myapp2",
	"Address": "myapp2",
	"Port": 10012,
	"check": {
		"id": "ping",
		"name": "HTTP api port on 10012",
		"http": "http://myapp2:10012/ping",
		"interval": "5s",
		"timeout": "1s"
	}
}
```





