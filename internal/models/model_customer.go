package models

// (cody) A Customer represents a shared model framework of the database and service models.
// often there is a segregation between db and service level to have different functions in each
type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"emailAddress"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}
