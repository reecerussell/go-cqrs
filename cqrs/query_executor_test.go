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

func TestExecute_GivenValidQuery_ReturnsResultWithValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	val := "Hello World"

	q := mock.NewMockQuery(ctrl)
	q.EXPECT().Validate().Return(nil)
	q.EXPECT().Execute(ctx).Return(val, nil)

	exe := cqrs.NewQueryExecutor()
	res := exe.Execute(ctx, q)

	assert.NoError(t, res.Err())
	assert.Equal(t, val, res.Value())
}

func TestExecute_GivenQueryWhereValidateFails_ReturnsResultWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	err := errors.New("test error")

	q := mock.NewMockQuery(ctrl)
	q.EXPECT().Validate().Return(err)

	exe := cqrs.NewQueryExecutor()
	res := exe.Execute(ctx, q)

	assert.Equal(t, err, res.Err())
	assert.Nil(t, res.Value())
}

func TestExecute_GivenQueryWhereExecuteFails_ReturnsResultWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	err := errors.New("test error")

	q := mock.NewMockQuery(ctrl)
	q.EXPECT().Validate().Return(nil)
	q.EXPECT().Execute(ctx).Return(nil, err)

	exe := cqrs.NewQueryExecutor()
	res := exe.Execute(ctx, q)

	assert.Equal(t, err, res.Err())
	assert.Nil(t, res.Value())
}
