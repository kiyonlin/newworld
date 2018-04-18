package scheduler

import "github.com/kiyonlin/newworld/crawler/engine"

// SimpleScheduler has worker channel
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// ConfigureMasterWorkerChan setup
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

// Submit a request to simple scheduler
func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() {
		s.workerChan <- request
	}()
}
