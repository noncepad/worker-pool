package manager

import (
	"context"

	"github.com/noncepad/worker-pool/pool"
	log "github.com/sirupsen/logrus"
)

func (e1 external[T, S]) Add(w pool.Worker[T, S]) error {
	deleteC := e1.deleteC
	signalC := w.CloseSignal()
	select {
	case <-e1.ctx.Done():
		return e1.ctx.Err()
	case e1.internalC <- func(in *internal[T, S]) {
		in.on_worker(w, signalC, deleteC)
	}:
		return nil
	}
}
func (in *internal[T, S]) on_worker(w pool.Worker[T, S], signalC <-chan error, deleteC chan<- deleteWorker) {
	in.workerIndex++
	id := in.workerIndex
	in.pool[id] = w
	go loopWorkerListen[T, S](in.ctx, in.jobC, deleteC, w, id)
}

type jobRequest[T, S any] struct {
	sendOnlyError bool
	job           pool.Job[T]
	respC         chan<- pool.Result[S]
	errorC        chan<- error
}

func loopWorkerListen[T, S any](
	ctx context.Context,
	jobC <-chan jobRequest[T, S],
	deleteC chan<- deleteWorker,
	w pool.Worker[T, S],
	id int,
) {
	log.Debugf("working %d listening for jobs", id)
	signalC := w.CloseSignal()
	doneC := ctx.Done()

	jobCount := 0
out:
	for {
		select {
		case <-doneC:
			break out
		case <-signalC:
			select {
			case <-doneC:
				break out
			case deleteC <- deleteWorker{id: id}:
			}
			break out
		case x := <-jobC:
			jobCount++
			log.Debugf("worker %d recieved job %d", id, jobCount)
			result, err := w.Run(x.job)
			if x.sendOnlyError {
				if err != nil {
					select {
					case <-doneC:
					case x.errorC <- err:
					}
				}
			} else {
				if err != nil {
					x.respC <- pool.ResultFromError[S](err)
				} else {
					x.respC <- pool.CreateResult[S](result)
				}

			}

		}
	}
	go w.Close()
}

func (in *internal[T, S]) on_delete(id int) {
	w, present := in.pool[id]
	if !present {
		return
	}
	delete(in.pool, id)
	go w.Close()
}
