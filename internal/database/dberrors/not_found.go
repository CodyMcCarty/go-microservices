package dberrors

import "fmt"

// NotFoundError (cody)
type NotFoundError struct {
	Entity string
	ID     string
}

// (cody)
// Creating a new instance of error so we don't leak database logic into our web layer
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("unable to find %s with id %s", e.Entity, e.ID)
}
