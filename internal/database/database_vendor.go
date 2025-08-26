package database

import (
	"context"
	"errors"

	"github.com/CodyMcCarty/go-microservices/internal/database/dberrors"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetAllVendors (cody) /database
func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Find(&vendors)
	return vendors, result.Error
}

// AddVendor (cody) /database
func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}

	return vendor, nil
}

func (c Client) GetVendorById(ctx context.Context, ID string) (*models.Vendor, error) {
	vendor := &models.Vendor{}
	result := c.DB.WithContext(ctx).
		Where(&models.Vendor{VendorID: ID}).
		First(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "vendor", ID: ID}
		}
		return nil, result.Error
	}
	return vendor, nil
}
