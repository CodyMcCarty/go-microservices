package server

import (
	"net/http"

	"github.com/CodyMcCarty/go-microservices/internal/database/dberrors"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

// GetAllVendors (cody) /server
func (s *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendors, err := s.DB.GetAllVendors(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

/* Post in chrome console:
fetch("http://localhost:8080/vendors", {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify({
    name: "John Doe Basic store",
    contact: "John Doe",
    phone: "(991) 321-6633",
    email: "jdoe@johndoebasic.com",
    address: "1005 Main St; Anytown, KS 66854"
  })
}).then(r => r.json())
  .then(data => console.log(data));

-- Post in bash --
curl -X POST http://localhost:8080/vendors \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Adv store",
    "contact": "John Doe",
    "phone": "(991) 321-6634",
    "email": "jdoe@johndoeadv.com",
    "address": "1006 Main St; Anytown, KS 66854"
  }'
*/

// AddVendor (cody) /server
func (s *EchoServer) AddVendor(ctx echo.Context) error {
	vendor := new(models.Vendor)
	if err := ctx.Bind(vendor); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	vendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, vendor)
}
