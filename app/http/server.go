package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mshirdel/quick/app"
	"github.com/mshirdel/quick/app/http/controller"
	"github.com/sirupsen/logrus"
)

type Server struct {
	app    *app.Application
	server *http.Server
}

func NewHTTPServer(app *app.Application) *Server {
	controller := controller.NewController(app)

	return &Server{
		app: app,
		server: &http.Server{
			Addr:         app.Cfg.Server.Address,
			ReadTimeout:  app.Cfg.Server.ReadTimeout,
			WriteTimeout: app.Cfg.Server.WriteTimeout,
			IdleTimeout:  app.Cfg.Server.IdleTimeout,
			Handler:      controller.Routes(),
		},
	}
}

func (s *Server) Start() {
	logrus.Infof("starting http server on: %s", s.app.Cfg.Server.Address)

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("failed starting http server: %v", err)
	}
}

func (s *Server) Shutdown() {
	logrus.Info("shutting down http server...")

	deadline, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(deadline); err != nil {
		logrus.Errorf("failed shutting down http server: %s", err)
	}
}
