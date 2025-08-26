package database

import (
	"context"

	"github.com/CodyMcCarty/go-microservices/internal/models"
)

// (cody) /database on Client model and added to DatabaseClient interface.
// always pass context. it allows you to write handlers in gorm to do specific things, such as if a user id is present, we can update the updatedBy
// where filter is conditional. if no email is passed, it won't use it
// is on client, pass ctx & email, return slice cust & err
func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customers)
	return customers, result.Error
}
