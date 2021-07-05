package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
	"time"
)

type Cfg struct {
	Port            int           `env:"SERVER_PORT" default:"8888"`
	ShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" default:"30s"`
}

type Server struct {
	cfg *Cfg
	log zerolog.Logger
	mx  *metrics

	eng *gin.Engine

	httpServ *http.Server
	mxHandler http.Handler
}

func New(cfg *Cfg, log zerolog.Logger, mxReg *prometheus.Registry) (*Server, error) {
	s := &Server{
		cfg: cfg,
		log: log,
		mx:  newMetrics(mxReg),
	}

	s.eng = gin.New()
	s.eng.Use(gin.Recovery())
	s.eng.Use(s.logRequest)

	s.httpServ = &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Port),
		Handler: s.eng,
	}

	s.mxHandler = promhttp.HandlerFor(mxReg, promhttp.HandlerOpts{})
	s.eng.GET("/metrics", s.getMetrics)

	go func() {
		if err := s.httpServ.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error().Err(err).Msg("ListenAndServe() failed")
		}
	}()
	s.log.Info().Msg("server started, check /ready for readiness")

	return s, nil
}

func (s *Server) Engine() *gin.Engine {
	return s.eng;
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()

	s.log.Info().Msg("shutting down")

	if err := s.httpServ.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful http server Shutdown() failed: %w", err)
	}
	s.log.Info().Msg("graceful shutdown complete")
	return nil
}

func (s *Server) logRequest(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	dur := time.Now().Sub(start)

	s.log.Info().
		Str("method", ctx.Request.Method).
		Str("uri", ctx.Request.RequestURI).
		Int("status", ctx.Writer.Status()).
		Dur("dur", dur).
		Msg("request processed")
}

func (s *Server) getMetrics(ctx *gin.Context) {
	s.mxHandler.ServeHTTP(ctx.Writer, ctx.Request)
}
