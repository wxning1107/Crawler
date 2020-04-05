package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request // in
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	// 解决循环等待问题，开goroutine
	go func() {
		s.workerChan <- r
	}()
}
