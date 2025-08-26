package database

import (
	"context"
	"errors"

	"github.com/CodyMcCarty/go-microservices/internal/database/dberrors"
	"github.com/CodyMcCarty/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetAllServices (cody) /database
func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Find(&services)
	return services, result.Error
}

// AddService (cody) /database
func (c Client) AddService(ctx context.Context, service *models.Service) (*models.Service, error) {
	service.ServiceID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}

	return service, nil
}
