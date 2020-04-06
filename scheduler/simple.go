package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	// workerChan is in
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	// 解决循环等待问题，开goroutine
	go func() {
		s.workerChan <- r
	}()
}
