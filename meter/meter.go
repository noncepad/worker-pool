package meter

import (
	"sync"

	pbt "github.com/noncepad/worker-pool/proto/solpipe"
	"google.golang.org/grpc"
)

type Hook struct {
	pbt.UnimplementedWorkerStatusServer
	busyWorkers  int
	totalWorkers int
	m            *sync.Mutex
}

func (h *Hook) OnStatus(req *pbt.Empty, stream pbt.WorkerStatus_OnStatusServer) error {
	return nil
}

func Create(s *grpc.Server) *Hook {

	hook := &Hook{
		busyWorkers:  0,
		totalWorkers: 0,
		m:            &sync.Mutex{},
	}
	pbt.RegisterWorkerStatusServer(s, hook)
	return hook
}

func (hook *Hook) JobStart() {
	hook.m.Lock()
	defer hook.m.Unlock()
	hook.busyWorkers++
}

func (hook *Hook) JobFinish() {
	hook.m.Lock()
	defer hook.m.Unlock()
	hook.busyWorkers--
}

func (hook *Hook) AddWorker() {
	hook.m.Lock()
	defer hook.m.Unlock()
	hook.totalWorkers++
}

func (hook *Hook) RemoveWorker() {
	hook.m.Lock()
	defer hook.m.Unlock()
	hook.totalWorkers--

}
