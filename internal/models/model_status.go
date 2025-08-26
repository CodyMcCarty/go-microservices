package models

// Health (cody)
// note: he often does a segregation between his db and service level models so he can have different functions in each. but here we use a shared model framework. 1. WireService 0:20
type Health struct {
	Status string `json:"status"`
}
