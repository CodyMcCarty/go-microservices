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
// different things may be errors.
// may need additional validation logic, those would need to be injected.
// if copy from data domain to web model, that may also modify this.
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

/*
http://localhost:8080/customers/f4a8473b-23cf-4084-83e3-7fd03ac43e36
*/

// GetCustomerById (cody) should I use the word 'Get', I've seen some go code omit 'Get' and just say CustomerById.
// -- In some c++ game code we use different things for a simple get that returns a cached or easily computed value vs something that requires computation.
// I'm a little confused when they use ptr vs data, ie *error in this case.
// NotFound only works when changing the ID, but the len remains the same. "record not found". should it look more like the other errors (see next comment)?
// if I put in !len, the console prints out the error "ERROR: invalid input syntax for type uuid: "f4a8473b-23cf-4084-83e3-7fd03ac43e3" (SQLSTATE 22P02)"
func (s *EchoServer) GetCustomerById(ctx echo.Context) error {
	ID := ctx.Param("id")
	customer, err := s.DB.GetCustomerById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, customer)
}
