package models

// (cody) shared model framework for the database and service layer
type Service struct {
	ServiceID string  `gorm:"primaryKey" json:"serviceId"`
	Name      string  `json:"name"`
	Price     float32 `gorm:"type:numeric" json:"price"`
}
