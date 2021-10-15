package cqrs

// ExecuteResult is the result of executing a command. Containing
// an error value and a result value of a command.
type ExecuteResult struct {
	err   error
	value interface{}
}

// NewExecuteResult returns a new instance of ExecuteResult.
func NewExecuteResult(err error, value interface{}) *ExecuteResult {
	return &ExecuteResult{
		err:   err,
		value: value,
	}
}

// Err returns a non-nil value if the resulting command
// failed to execute.
func (r *ExecuteResult) Err() error {
	return r.err
}

// Value returns the value of a CommandWithValue command.
// If the resulting command was not of type CommandWithValue,
// the value will always be nil.
func (r *ExecuteResult) Value() interface{} {
	return r.value
}
