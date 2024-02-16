package manager_test

import (
	"context"
	"testing"

	"github.com/noncepad/worker-pool/manager"
	"github.com/noncepad/worker-pool/pool"
	log "github.com/sirupsen/logrus"
)

func TestBasic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mgr, err := manager.Create[int, int](ctx, 10)
	if err != nil {
		t.Fatal(err)
	}
	signalC := mgr.CloseSignal()
	t.Cleanup(func() {
		<-signalC
		log.Info("manager exited")
	})
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		err = mgr.Add(createAdder(ctx))
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < 1_000; i++ {
		input := i
		ans, err := mgr.Submit(ctx, input)
		if err != nil {
			t.Fatal(err)
		}
		if ans != myAdder(input) {
			t.Fatalf("got wrong output: %d vs %d", ans, myAdder(input))
		}
	}

}

func myAdder(a int) int {
	return a + 1
}

type adderWorker struct {
	ctx        context.Context
	cancel     context.CancelFunc
	reqSignalC chan<- chan<- error
}

func createAdder(parentCtx context.Context) *adderWorker {
	reqSignalC := make(chan chan<- error, 1)
	ctx, cancel := context.WithCancel(parentCtx)
	go loopWait(ctx, cancel, reqSignalC)
	return &adderWorker{ctx: ctx, cancel: cancel, reqSignalC: reqSignalC}
}

func loopWait(ctx context.Context, cancel context.CancelFunc, reqSignalC <-chan chan<- error) {
	doneC := ctx.Done()
	closeCList := make([]chan<- error, 0)
out:
	for {
		select {
		case <-doneC:
			break out
		case errorC := <-reqSignalC:
			closeCList = append(closeCList, errorC)
		}
	}
	for _, errorC := range closeCList {
		errorC <- nil
	}
}

func (aw *adderWorker) Run(job pool.Job[int]) (int, error) {
	return myAdder(job.Payload()), nil
}

func (aw *adderWorker) Close() error {
	signalC := aw.CloseSignal()
	aw.cancel()
	return <-signalC
}

func (aw *adderWorker) CloseSignal() <-chan error {
	signalC := make(chan error, 1)
	select {
	case <-aw.ctx.Done():
		signalC <- aw.ctx.Err()
	case aw.reqSignalC <- signalC:
	}
	return signalC
}
