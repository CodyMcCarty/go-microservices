package database

import (
	"context"

	"github.com/CodyMcCarty/go-microservices/internal/models"
)

// (cody) /database
func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Find(&vendors)
	return vendors, result.Error
}
