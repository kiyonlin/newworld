package scheduler

import "github.com/kiyonlin/newworld/crawler/engine"

// SimpleScheduler has worker channel
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// WorkerChan returns a request channel
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// WorkerReady knows if worker is ready
func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

// Run init simple scheduler's worker channel
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// Submit a request to simple scheduler
func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() {
		s.workerChan <- request
	}()
}
