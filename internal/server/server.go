package server

import (
	"log"
	"net/http"

	"github.com/CodyMcCarty/go-microservices/internal/database"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

// Server (cody)
// Server interface & EchoServer struct
// Readiness and Liveness checks like in kubernetes
type Server interface {
	// Start (cody) in Server interface
	Start() error

	// Readiness (cody) in Server interface
	Readiness(ctx echo.Context) error

	// Liveness (cody) in Server interface
	Liveness(ctx echo.Context) error

	// GetAllCustomers (cody) in Server interface
	GetAllCustomers(ctx echo.Context) error
	AddCustomer(ctx echo.Context) error
	GetCustomerById(ctx echo.Context) error
	UpdateCustomer(ctx echo.Context) error
	DeleteCustomer(ctx echo.Context) error

	// GetAllProducts (cody) in Server interface
	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error

	// GetAllServices (cody) in Server interface
	GetAllServices(ctx echo.Context) error
	AddService(ctx echo.Context) error
	GetServiceById(ctx echo.Context) error

	// GetAllVendors (cody) in Server interface
	GetAllVendors(ctx echo.Context) error
	AddVendor(ctx echo.Context) error
	GetVendorById(ctx echo.Context) error
}

// EchoServer (cody)
// Server interface & EchoServer struct
type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

// (cody)
// creates a customer group. not necessary and doesn't make sense for simple method. 2.GetAllOp 5:40. As build more methods, becomes more useful especially if passing middleware. we won't have any middleware
// at customers with nothing but the prefix a GET op will call GetAllCustomers.
// /customers?emailAddress=nibh@ultricesposuere.edu . let's do breakpoints and step through code.
// products?vendorId=87e31b27-7f6f-41b7-a3f0-3d2a4fcd67e9
func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/customers")
	cg.GET("", s.GetAllCustomers)
	cg.POST("", s.AddCustomer)
	cg.GET("/:id", s.GetCustomerById)
	cg.PUT("/:id", s.UpdateCustomer)
	cg.DELETE("/:id", s.DeleteCustomer)

	pg := s.echo.Group("/products")
	pg.GET("", s.GetAllProducts)
	pg.POST("", s.AddProduct)
	pg.GET("/:id", s.GetProductById)

	sg := s.echo.Group("/services")
	sg.GET("", s.GetAllServices)
	sg.POST("", s.AddService)
	sg.GET("/:id", s.GetServiceById)

	vg := s.echo.Group("/vendors")
	vg.GET("", s.GetAllVendors)
	vg.POST("", s.AddVendor)
	vg.GET("/:id", s.GetVendorById)
}

// NewEchoServer (cody)
// he set it up this way for testing 1.SetUpEchoClient 1:27. We're not testing in this course. Can inject a mock
func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

// (cody)
// we should: port 8080 from config 1.setupEchoClient 2:50
func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred %s", err)
		return err
	}
	return nil
}

// Readiness (cody)
func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusServiceUnavailable, models.Health{Status: "Failure"})
}

// Liveness (cody)
func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
