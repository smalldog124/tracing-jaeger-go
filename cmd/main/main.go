package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	jaegerClient "github.com/uber/jaeger-client-go"
)

func main() {
	serviceName := "smalldog"
	udpSender, err := jaegerClient.NewUDPTransport("localhost:6831", 1024)
	if err != nil {
		log.Fatal("udp sender fail", err)
	}
	reporter := jaegerClient.NewRemoteReporter(udpSender)
	sampler := jaegerClient.NewRateLimitingSampler(1)

	tracer, closer := jaegerClient.NewTracer(
		serviceName,
		sampler,
		reporter,
	)
	defer closer.Close()

	route := gin.Default()
	route.GET("/hello", func(context *gin.Context) {
		url := fmt.Sprintf("[%s] %s%s", context.Request.Method, context.Request.Host, context.Request.URL)
		span := tracer.StartSpan("http_server")
		defer span.Finish()
		span.SetTag("http.url", url)
		span.SetTag("http.status_code", http.StatusOK)
		log.Println("hello!!!")
		context.JSON(http.StatusOK, gin.H{"masses": "hello"})
	})
	route.Run(":8880")
}
