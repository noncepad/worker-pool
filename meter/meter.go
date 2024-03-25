package meter

import (
	"context"
	"sync"
	"time"

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
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// Loop to send a response every time the ticker ticks
	for {
		select {
		case <-ctx.Done():
			// stream closed
			return nil
		case <-ticker.C:
			// Send a response
			h.m.Lock()
			capacity := float32(h.busyWorkers) / float32(h.totalWorkers)
			h.m.Unlock()
			if err := stream.Send(&pbt.CapacityResponse{Capacity: capacity}); err != nil {
				return err
			}
		}
	}
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

func (h *Hook) JobStart() {
	h.m.Lock()
	defer h.m.Unlock()
	h.busyWorkers++
}

func (h *Hook) JobFinish() {
	h.m.Lock()
	defer h.m.Unlock()
	h.busyWorkers--
}

func (h *Hook) AddWorker() {
	h.m.Lock()
	defer h.m.Unlock()
	h.totalWorkers++
}

func (h *Hook) RemoveWorker() {
	h.m.Lock()
	defer h.m.Unlock()
	h.totalWorkers--
}
