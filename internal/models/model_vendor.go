package models

// (cody) shared model framework for the database and service layer
type Vendor struct {
	VendorID string `gorm:"primaryKey" json:"vendorId"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
