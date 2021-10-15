package cqrs

import "context"

// CommandExecutor is used to execute Commands.
type CommandExecutor interface {
	// Execute executes cmd, with the given context. Before execution
	// the command is validated using cmd.Validate(). An instance of
	// ExecuteResult will be returned, always containing a non-nil value.
	Execute(ctx context.Context, cmd CommandBase) *ExecuteResult
}

type commandExecutor struct{}

// NewCommandExecutor returns a new instance of CommandExecutor.
func NewCommandExecutor() CommandExecutor {
	return &commandExecutor{}
}

func (*commandExecutor) Execute(ctx context.Context, cmd CommandBase) *ExecuteResult {
	err := cmd.Validate()
	if err != nil {
		return NewExecuteResult(err, nil)
	}

	switch t := cmd.(type) {
	case CommandWithValue:
		v, err := t.Execute(ctx)
		return NewExecuteResult(err, v)
	default:
		c := cmd.(Command)
		err = c.Execute(ctx)
		return NewExecuteResult(err, nil)
	}
}
