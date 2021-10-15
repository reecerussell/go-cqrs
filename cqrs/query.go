package cqrs

import "context"

// Query is used to perform read operations.
type Query interface {
	// Execute is used to execute the query.
	Execute(ctx context.Context) (interface{}, error)

	// Validate is used to ensure that any data passed to a query
	// is valid, before execution. If an error is returned, the
	// query is not valid.
	Validate() error
}
