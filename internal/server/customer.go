package server

import (
	"net/http"

	"github.com/CodyMcCarty/go-microservices/internal/database/dberrors"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

// GetAllCustomers (cody) /server
// we pass echo.Context instead of context.Context
func (s *EchoServer) GetAllCustomers(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("emailAddress")

	customers, err := s.DB.GetAllCustomers(ctx.Request().Context(), emailAddress)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, customers)
}

/* Post in chrome console:
fetch("http://localhost:8080/customers", {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify({
    firstName: "John",
    lastName: "Doe",
    emailAddress: "jdoe1@example.com",
    phoneNumber: "515-555-1234",
    address: "1234 Main St; Anytown, KS 66854"
  })
}).then(r => r.json())
  .then(data => console.log(data));

-- Post in bash --
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "emailAddress": "jdoe@example.com",
    "phoneNumber": "515-555-1234",
    "address": "1234 Main St; Anytown, KS 66854"
  }'
*/

// AddCustomer (cody) /server
// new to get a ptr.
// I think Bind() sets headers and stuff.
// we could validate that it is StatusUnsupportedMediaType.
func (s *EchoServer) AddCustomer(ctx echo.Context) error {
	customer := new(models.Customer)
	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	customer, err := s.DB.AddCustomer(ctx.Request().Context(), customer)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, customer)
}
