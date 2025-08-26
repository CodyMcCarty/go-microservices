package database

import (
	"context"

	"github.com/CodyMcCarty/go-microservices/internal/models"
)

// (cody) /database
func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorId}).
		Find(&products)
	return products, result.Error
}
