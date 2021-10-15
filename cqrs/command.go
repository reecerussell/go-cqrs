package cqrs

import "context"

// CommandBase is used to provide core funcs and properties to a command.
type CommandBase interface {
	// Validate is used to ensure that any data passed to a command
	// is valid, before execution. If an error is returned, the
	// command is not valid.
	Validate() error
}

// Command is an interface used to perform a write operation or perform an action.
type Command interface {
	CommandBase

	// Execute is used to run the command.
	Execute(ctx context.Context) error
}

// Command is an interface used to perform a write operation or perform an action,
// similar to Command, however, on execute a value is returned.
type CommandWithValue interface {
	CommandBase

	// Execute is used to run the command.
	Execute(ctx context.Context) (interface{}, error)
}
