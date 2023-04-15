package http

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"onelab/internal/config"
	"onelab/internal/jsonlog"
	"onelab/internal/transport/http/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Manager
	App     *echo.Echo
}

func NewServer(cfg *config.Config, handler *handler.Manager) *Server {
	return &Server{
		cfg:     cfg,
		handler: handler,
	}
}

func (s *Server) StartServer(ctx context.Context, logger *jsonlog.Logger) error {
	s.App = s.BuildEngine(logger)
	s.NewRouter()
	shotdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		q := <-quit

		logger.PrintInfo("shutting down server signal:", map[string]string{
			"signal": q.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shotdownError <- s.App.Shutdown(ctx)
	}()

	logger.PrintInfo("starting server", map[string]string{
		"addr": s.cfg.HTTP.HttpPort,
	})

	err := s.App.Start(s.cfg.HTTP.HttpPort)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shotdownError
	if err != nil {
		return err
	}

	logger.PrintInfo("stopped server", map[string]string{
		"addr": s.cfg.HTTP.HttpPort,
	})

	return nil
}

func (s *Server) BuildEngine(l *jsonlog.Logger) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l.PrintInfo("request", map[string]string{
				"URI": v.URI,
			})
			return nil
		},
	}))

	return e
}
