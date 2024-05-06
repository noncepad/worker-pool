package meter

import (
	"context"
	"sync"
	"time"

	pbt "github.com/noncepad/solpipe-market/go/proto/relay"
	"google.golang.org/grpc"
)

type Hook struct {
	pbt.UnimplementedCapacityServer
	busyWorkers  int
	totalWorkers int
	m            *sync.Mutex
	subIndex     int
	subM         map[int]*hookSub
}

type hookSub struct {
	ctx    context.Context
	cancel context.CancelFunc
	i      int
	outC   chan<- float32
}

func (h *Hook) calculateCapacity() float32 {
	h.m.Lock()
	defer h.m.Unlock()
	return h.unsafeCalculateCapacity()
}

func (h *Hook) unsafeCalculateCapacity() float32 {

	a := h.busyWorkers
	b := h.totalWorkers

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

func (h *Hook) OnStatus(req *pbt.Empty, stream pbt.Capacity_OnStatusServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	h.m.Lock()
	x := new(hookSub)
	x.i = h.subIndex
	h.subIndex++
	changeC := make(chan float32, 100)
	x.outC = changeC
	x.ctx = ctx
	x.cancel = cancel
	h.subM[x.i] = x
	h.m.Unlock()

	var err error
	// Loop to send a response every time the ticker ticks
out:
	for {
		select {
		case <-ctx.Done():
			// stream closed
			break out
		case capacity := <-changeC:
			if 1 < capacity {
				capacity = 1
			} else if capacity < 0 {
				capacity = 0
			}
			if err = stream.Send(&pbt.CapacityStatus{UtilizationRatio: capacity}); err != nil {
				break out
			}
		case <-ticker.C:
			// Send a response
			capacity := h.calculateCapacity()
			if 1 < capacity {
				capacity = 1
			} else if capacity < 0 {
				capacity = 0
			}
			if err = stream.Send(&pbt.CapacityStatus{UtilizationRatio: capacity}); err != nil {
				break out
			}
		}
	}
	h.m.Lock()
	delete(h.subM, x.i)
	h.m.Unlock()
	return err
}
func Create(s *grpc.Server) *Hook {

	hook := &Hook{
		busyWorkers:  0,
		totalWorkers: 0,
		m:            &sync.Mutex{},
		subIndex:     0,
		subM:         make(map[int]*hookSub),
	}
	pbt.RegisterCapacityServer(s, hook)
	return hook
}

func (h *Hook) JobStart() {
	h.m.Lock()
	defer h.m.Unlock()
	h.busyWorkers++
	h.unsafeCapacityUpdate()
}

func (h *Hook) JobFinish() {
	h.m.Lock()
	defer h.m.Unlock()
	h.busyWorkers--
	h.unsafeCapacityUpdate()
}

func (h *Hook) WorkerAdd() {
	h.m.Lock()
	defer h.m.Unlock()
	h.totalWorkers++
	h.unsafeCapacityUpdate()
}

func (h *Hook) WorkerRemove() {
	h.m.Lock()
	defer h.m.Unlock()
	h.totalWorkers--
	h.unsafeCapacityUpdate()
}

func (h *Hook) unsafeCapacityUpdate() {
	capacity := h.unsafeCalculateCapacity()
	for _, v := range h.subM {
		select {
		case v.outC <- capacity:
		default:
			v.cancel()
			delete(h.subM, v.i)
		}
	}
}
