# Jaeger tracing your services

## Steps

### Run Tracing server(Jaeger)   on Docker
```
$docker container run -d -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 jaegertracing/all-in-one:latest

```
Open Jaeger UI at http://localhost:16686/

### Run your service

1. Run service `go run cmd/main/main.go`
2. Call http://localhost:8880/hello
3. See detail of tracing service at http://localhost:16686/