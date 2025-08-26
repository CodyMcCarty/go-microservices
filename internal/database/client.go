package database

import (
	"context"
	"fmt"
	"time"

	"github.com/CodyMcCarty/go-microservices/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DatabaseClient (cody)
// we don't use DatabaseClient interface, we use Client struct?
type DatabaseClient interface {
	// Ready (cody) in DatabaseClient interface
	Ready() bool

	// GetAllCustomers (cody) in DatabaseClient interface
	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)

	// GetAllProducts (cody) in DatabaseClient interface
	GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)

	// GetAllServices (cody) in DatabaseClient interface
	GetAllServices(ctx context.Context) ([]models.Service, error)

	// GetAllVendors (cody) in DatabaseClient interface
	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
}

// Client (cody)
// 1:44 set up db
// we don't use DatabaseClient interface, we use Client struct?
type Client struct {
	DB *gorm.DB
}

// NewDatabaseClient (cody)
// we should pass in config. Frank says he typically passes in the config, but not for this course. 1.SetUpTheDatabaseClient 1:40
// the dsn host is localhost, but should have the ability to be configured as something else like the remote instance.
// TablePrefix: bc he defined everything in a schema, we need to add this or reference it everywhere and that's too painful 1.SetUpDBClient 4:02
func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		5432,
		"disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time { // why do I need this function?
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return client, nil
}

// Ready (cody)
// tx = transaction
func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
