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
		in.on_worker(w, signalC, deleteC, e1.hook)
	}:
		return nil
	}
}
func (in *internal[T, S]) on_worker(w pool.Worker[T, S], signalC <-chan error, deleteC chan<- deleteWorker, hook JobHook) {
	in.workerIndex++
	id := in.workerIndex
	in.pool[id] = w
	go loopWorkerListen[T, S](in.ctx, in.jobC, deleteC, signalC, w, id, hook)
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
	signalC <-chan error,
	w pool.Worker[T, S],
	id int,
	hook JobHook,
) {
	log.Debugf("working %d listening for jobs", id)
	doneC := ctx.Done()

	jobCount := 0
out:
	for {
		//log.Debugf("worker %d - 0 - recieved job %d", id, jobCount)
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
			//log.Debugf("worker %d - 1 - recieved job %d", id, jobCount)
			// Add hook 1: Worker takes a job (becomes busy)
			hook.JobStart()
			result, err := w.Run(x.job)
			// Add hook 2: Worker takes a job (becomes idle)
			hook.JobFinish()
			if x.sendOnlyError {
				if err != nil {
					//log.Debugf("worker %d - 2 - recieved job %d", id, jobCount)
					select {
					case <-doneC:
					case x.errorC <- err:
						//log.Debugf("worker %d - 3 - recieved job %d", id, jobCount)
					}
				}
			} else {
				if err != nil {
					//log.Debugf("worker %d - 5 - recieved job %d", id, jobCount)
					x.respC <- pool.ResultFromError[S](err)
				} else {
					//log.Debugf("worker %d - 6 - recieved job %d", id, jobCount)
					x.respC <- pool.CreateResult[S](result)
				}
				//log.Debugf("worker %d - 7 - recieved job %d", id, jobCount)

			}

		}
	}
	log.Debugf("worker %d exiting loop", id)
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
