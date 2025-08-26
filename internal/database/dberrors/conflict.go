package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	// note: creating a new instance of error so we don't leak database logic into our web layer
	return "attempted to create a record with an existing key"
}
