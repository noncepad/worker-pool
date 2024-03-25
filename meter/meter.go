package meter

import "sync"

type Hook struct {
	busyWorkers  int
	totalWorkers int
	m            *sync.Mutex
}

func Create() *Hook {
	return &Hook{
		busyWorkers:  0,
		totalWorkers: 0,
		m:            &sync.Mutex{},
	}
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
