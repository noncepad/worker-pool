package manager

import (
	"context"

	"github.com/noncepad/worker-pool/pool"
	log "github.com/sirupsen/logrus"
)

type internal[T, S any] struct {
	ctx              context.Context
	closeSignalCList []chan<- error
	hook             JobHook
	jobC             <-chan jobRequest[T, S]
	workerIndex      int
	pool             map[int]pool.Worker[T, S]
}

func loopInternal[T, S any](
	ctx context.Context,
	cancel context.CancelFunc,
	hook JobHook,
	internalC <-chan func(*internal[T, S]),
	deleteC <-chan deleteWorker,
	jobC <-chan jobRequest[T, S],
) {
	defer cancel()
	doneC := ctx.Done()
	var err error

	in := new(internal[T, S])
	in.ctx = ctx
	in.closeSignalCList = make([]chan<- error, 0)
	in.hook = hook
	in.jobC = jobC
	in.pool = make(map[int]pool.Worker[T, S])

out:
	for {
		select {
		case <-doneC:
			break out
		case req := <-internalC:
			req(in)
		case x := <-deleteC:
			in.on_delete(x.id)
		}
	}
	in.finish(err)
}

func (in *internal[T, S]) finish(err error) {
	log.Debugf("exiting manager pool: %s", err)
	for _, errorC := range in.closeSignalCList {
		errorC <- err
	}
}
