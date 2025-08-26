package server

import (
	"net/http"

	"github.com/CodyMcCarty/go-microservices/internal/database/dberrors"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

// GetAllServices (cody) /server
func (s *EchoServer) GetAllServices(ctx echo.Context) error {
	services, err := s.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, services)
}

/* Post in chrome console:
fetch("http://localhost:8080/services", {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify({
    name: "John Doe cleaning",
    price: 49.95
  })
}).then(r => r.json())
  .then(data => console.log(data));

-- Post in bash --
curl -X POST http://localhost:8080/services \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe XL cleaning",
    "price": 74.95
  }'
*/

// AddService (cody) /server
func (s *EchoServer) AddService(ctx echo.Context) error {
	service := new(models.Service)
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	service, err := s.DB.AddService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, service)
}

func (s *EchoServer) GetServiceById(ctx echo.Context) error {
	ID := ctx.Param("id")
	service, err := s.DB.GetServiceById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, service)
}
