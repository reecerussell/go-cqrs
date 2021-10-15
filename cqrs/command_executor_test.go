package cqrs_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/reecerussell/go-cqrs/cqrs"
	"github.com/reecerussell/go-cqrs/cqrs/mock"
)

func TestExecute_GivenValidCommand_ExecutesSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	cmd := mock.NewMockCommand(ctrl)
	cmd.EXPECT().Validate().Return(nil)
	cmd.EXPECT().Execute(ctx).Return(nil)

	exe := cqrs.NewCommandExecutor()
	res := exe.Execute(ctx, cmd)

	assert.NoError(t, res.Err())
	assert.Nil(t, res.Value())
}

func TestExecute_GivenCommandWhereValidateFails_ResultWithErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	err := errors.New("test error")

	cmd := mock.NewMockCommand(ctrl)
	cmd.EXPECT().Validate().Return(err)

	exe := cqrs.NewCommandExecutor()
	res := exe.Execute(ctx, cmd)

	assert.Equal(t, err, res.Err())
	assert.Nil(t, res.Value())
}

func TestExecute_GivenCommandWhereExecuteFails_ResultWithErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	err := errors.New("test error")

	cmd := mock.NewMockCommand(ctrl)
	cmd.EXPECT().Validate().Return(nil)
	cmd.EXPECT().Execute(ctx).Return(err)

	exe := cqrs.NewCommandExecutor()
	res := exe.Execute(ctx, cmd)

	assert.Equal(t, err, res.Err())
	assert.Nil(t, res.Value())
}

func TestExecute_GivenValidCommandWithValue_ExecutesSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	cmd := mock.NewMockCommandWithValue(ctrl)
	cmd.EXPECT().Validate().Return(nil)
	cmd.EXPECT().Execute(ctx).Return("Hello World", nil)

	exe := cqrs.NewCommandExecutor()
	res := exe.Execute(ctx, cmd)

	assert.NoError(t, res.Err())
	assert.Equal(t, "Hello World", res.Value())
}

func TestExecute_GivenCommandWithValueWhereValidateFails_ResultWithErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	err := errors.New("test: error")

	cmd := mock.NewMockCommandWithValue(ctrl)
	cmd.EXPECT().Validate().Return(err)

	exe := cqrs.NewCommandExecutor()
	res := exe.Execute(ctx, cmd)

	assert.Equal(t, err, res.Err())
	assert.Nil(t, res.Value())
}

func TestExecute_GivenCommandWithValueWhereExecuteFails_ResultWithErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	err := errors.New("test: error")

	cmd := mock.NewMockCommandWithValue(ctrl)
	cmd.EXPECT().Validate().Return(nil)
	cmd.EXPECT().Execute(ctx).Return(nil, err)

	exe := cqrs.NewCommandExecutor()
	res := exe.Execute(ctx, cmd)

	assert.Equal(t, err, res.Err())
	assert.Nil(t, res.Value())
}
