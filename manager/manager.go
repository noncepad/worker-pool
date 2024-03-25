package manager

import (
	"context"

	"github.com/noncepad/worker-pool/pool"
)

type external[T, S any] struct {
	ctx       context.Context
	cancel    context.CancelFunc
	internalC chan<- func(*internal[T, S])
	deleteC   chan<- deleteWorker
	jobC      chan<- jobRequest[T, S]
	hook      JobHook
}

type JobHook interface {
	JobStart()
	JobFinish()
	AddWorker()
	RemoveWorker()
}

type deleteWorker struct {
	id int
}

func Create[T, S any](
	parentCtx context.Context,
	bufferSize int,
	hook JobHook,
) (pool.Manager[T, S], error) {
	ctx, cancel := context.WithCancel(parentCtx)
	internalC := make(chan func(*internal[T, S]), 10)
	deleteC := make(chan deleteWorker, 10)
	jobC := make(chan jobRequest[T, S], bufferSize)
	go loopInternal[T, S](
		ctx,
		cancel,
		internalC,
		deleteC,
		jobC,
	)
	return external[T, S]{
		ctx:       ctx,
		cancel:    cancel,
		internalC: internalC,
		jobC:      jobC,
	}, nil
}

func (e1 external[T, S]) Close() error {
	signalC := e1.CloseSignal()
	e1.cancel()
	return <-signalC
}

func (e1 external[T, S]) CloseSignal() <-chan error {
	signalC := make(chan error, 1)
	select {
	case <-e1.ctx.Done():
		signalC <- e1.ctx.Err()
	case e1.internalC <- func(in *internal[T, S]) {
		in.closeSignalCList = append(in.closeSignalCList, signalC)
	}:
	}
	return signalC
}

func (e1 external[T, S]) DetachSubmit(
	ctx context.Context,
	payload T,
	id int,
	errorC chan<- error,
) {

	doneC := e1.ctx.Done()
	cancel, job := pool.CreateJob[T](ctx, payload)
	defer cancel()
	select {
	case <-doneC:
		return
	case e1.jobC <- jobRequest[T, S]{sendOnlyError: true, job: job, errorC: errorC}:
	}
}

func (e1 external[T, S]) Submit(ctx context.Context, payload T) (result S, err error) {
	doneC := e1.ctx.Done()
	respC := make(chan pool.Result[S], 1)
	cancel, job := pool.CreateJob[T](ctx, payload)
	defer cancel()
	select {
	case <-doneC:
		err = e1.ctx.Err()
		return
	case e1.jobC <- jobRequest[T, S]{sendOnlyError: false, job: job, respC: respC}:
	}
	var resp pool.Result[S]
	select {
	case <-doneC:
		err = e1.ctx.Err()
		return
	case resp = <-respC:
	}
	err = resp.Err()
	if err != nil {
		return
	}
	result = resp.Payload()
	return
}
