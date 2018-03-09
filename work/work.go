// Package work manages a pool of goroutines to perform work.
package work

import (
	"sync"
)

// Worker must be implemented by types that want to use
// the work pool.
type Worker interface {
	Task()
}

// Pool provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type Pool struct {
	// 无缓冲通道保证调用的 Run 返回时，提交的工作已经开始执行
	work chan Worker
	wg   sync.WaitGroup
}

// New creates a new work pool.
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// 无缓冲通道保证调用的 Run 返回时，提交的工作已经开始执行
			// p.work 被close 时才会停止循环
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run sibmit work to the pool.
func (p *Pool) Run(w Worker) {
	// 无缓冲通道保证调用的 Run 返回时，提交的工作已经开始执行
	p.work <- w
}

// Shutdown waits for all the goroutines to shutdown.
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
