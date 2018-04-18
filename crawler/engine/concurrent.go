package engine

// ConcurrentEngine has scheduler and worker count
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

// Scheduler can sumbit and configure master worker channel
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// ReadyNotifier knows if worker is ready
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run all request seeds
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.URL) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(item interface{}) {
				e.ItemChan <- item
			}(item)
		}
		for _, request := range result.Requests {
			if isDuplicate(request.URL) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedURLs = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedURLs[url] {
		return true
	}
	visitedURLs[url] = true
	return false
}
