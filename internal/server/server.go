package server

import (
	"log"
	"net/http"

	"github.com/CodyMcCarty/go-microservices/internal/database"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	// Readiness and Liveness checks like in kubernetes
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	// he set it up this way for testing 1.SetUpEchoClient 1:27. We're not testing in this course. Can inject a mock
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	// todo port 8080 from config 1.setupEchoClient 2:50
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusServiceUnavailable, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
