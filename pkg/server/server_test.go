package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"gopkg.in/resty.v1"
	"testing"
	"time"
)

func TestServerSimple(t *testing.T) {
	s, err := New(&Cfg{
		Port:            9999,
		ShutdownTimeout: time.Second,
	}, zerolog.Nop(), prometheus.NewRegistry())
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	s.Engine().GET("/test", func(context *gin.Context) {
		context.Status(200)
	})

	res, err := resty.New().R().Get("http://localhost:9999/test")
	if err != nil {
		t.Fatal(err)
	}
	if s := res.StatusCode(); s != 200 {
		t.Fatalf("status: %d", s)
	}
}
