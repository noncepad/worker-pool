package pool

import (
	"context"
)

type Base interface {
	Close() error
	CloseSignal() <-chan error
}

type Manager[T, S any] interface {
	Base
	// add a worker to the pool
	Add(Worker[T, S]) error
	// submit a payload to be done
	Submit(context.Context, T) (S, error)
	// only send an error if there is an error
	DetachSubmit(context.Context, T, int, chan<- error)
}

type DetachSubmitResult[S any] struct {
	Err    error
	Result S
	Id     int
}

type Worker[T, S any] interface {
	Base
	Run(Job[T]) (S, error)
}

type Job[T any] struct {
	ctx     context.Context
	payload T
}

func CreateJob[T any](parentCtx context.Context, payload T) (context.CancelFunc, Job[T]) {
	ctx, cancel := context.WithCancel(parentCtx)
	return cancel, Job[T]{ctx: ctx, payload: payload}
}

func (j Job[T]) Ctx() context.Context {
	return j.ctx
}

func (j Job[T]) Payload() T {
	return j.payload
}

type Result[T any] struct {
	err     error
	payload T
}

func CreateResult[T any](payload T) Result[T] {
	return Result[T]{payload: payload}
}
func ResultFromError[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (j Result[T]) Err() error {
	return j.err
}

func (j Result[T]) Payload() T {
	return j.payload
}
