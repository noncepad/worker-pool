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

func (h *Hook) calculateCapacity() float32 {
	h.m.Lock()
	a := h.busyWorkers
	b := h.totalWorkers
	h.m.Unlock()

	capacity := float32(a) / float32(b)

	if b == 0 {
		return 1
	}

	if a < 0 || b < 0 {
		return 1
	}
	if b < a {
		return 1
	}
	return capacity
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
			capacity := h.calculateCapacity()
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
