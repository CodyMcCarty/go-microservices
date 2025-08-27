package database

import (
	"context"
	"errors"

	"github.com/CodyMcCarty/go-microservices/internal/database/dberrors"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetAllCustomers (cody) /database on Client model and added to DatabaseClient interface.
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

// AddCustomer (cody). takes and returns ptr to Customer.
// check for duplicated key, and should fix it, but we didn't.
func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return customer, nil
}

// GetCustomerById (cody) returns ptr
func (c Client) GetCustomerById(ctx context.Context, ID string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := c.DB.WithContext(ctx).
		Where(&models.Customer{CustomerID: ID}).
		First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "customer", ID: ID}
		}
		return nil, result.Error
	}
	return customer, nil
}

// UpdateCustomer (cody) takes ptr, returns ptr.
// update is something you have to spend time to think about.
// why slice of customers?
// we will guard against modifying the customerID in the web layer.
// emails are not unique in our db, but they should be
// Updates() doesn't have to include everything i.e. ID.
// why are the err switch err.(type) sometimes and if errors.Is other times?
// we return [0] to be sure that we get the obj that returns. 2.UpdateOp 4:30. what does that mean?
// the way he does it is broken. I bandaid it. Need to find propper way.
// what is allowed to be updated vs retaining some fields and not updating them. 2.ChUpdate 0:20 & what to use as a patch instead of a put. service only has a few fields instead of PUT do a PATCH.
func (c Client) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	var customers []models.Customer

	result := c.DB.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where(&models.Customer{CustomerID: customer.CustomerID}).
		Updates(&models.Customer{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Address:   customer.Address,
		}).
		Scan(&customers)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "customer", ID: customer.CustomerID}
	}

	return &customers[0], nil
}

func (c Client) DeleteCustomer(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).
		Delete(&models.Customer{CustomerID: ID}).
		Error
}
