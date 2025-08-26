package dberrors

// ConflictError (cody)
type ConflictError struct{}

// (cody)
// Creating a new instance of error so we don't leak database logic into our web layer
func (e *ConflictError) Error() string {
	return "attempted to create a record with an existing key"
}
