package cqrs

import "context"

// QueryExecutor is used to execute Querys.
type QueryExecutor interface {
	// Execute is used to execute a query. Always returning
	// a non-nil ExecuteResult, containing the query's result
	// and an error, if any.
	Execute(ctx context.Context, q Query) *ExecuteResult
}

type queryExecutor struct{}

// NewQueryExecutor returns a new instance of QueryExecutor.
func NewQueryExecutor() QueryExecutor {
	return &queryExecutor{}
}

func (*queryExecutor) Execute(ctx context.Context, q Query) *ExecuteResult {
	err := q.Validate()
	if err != nil {
		return NewExecuteResult(err, nil)
	}

	v, err := q.Execute(ctx)
	if err != nil {
		return NewExecuteResult(err, nil)
	}

	return NewExecuteResult(nil, v)
}
