package api

import (
	"sync"
)

type Jobfn func() error

type WorkerPool struct {
	jobCh       chan Jobfn
	jobResCh    chan error
	jobCancelCh chan bool
	wg          sync.WaitGroup
	concurrent  int
	maxErrors   int
	errors      []error
}

func (wp *WorkerPool) worker() {

	defer wp.wg.Done()

	for {
		select {
		case <-wp.jobCancelCh:
			return
		case job := <-wp.jobCh:

			if job == nil {
				return
			}

			err := job()
			wp.jobResCh <- err
		}
	}

}

func NewWorkerPool(jobs []Jobfn, N int, maxErrors int) WorkerPool {

	wp := WorkerPool{}

	wp.jobCh = make(chan Jobfn, len(jobs))
	wp.jobResCh = make(chan error, len(jobs))
	wp.jobCancelCh = make(chan bool, N)
	wp.errors = make([]error, 0)

	for _, job := range jobs {
		wp.jobCh <- job
	}

	wp.concurrent = N
	wp.maxErrors = maxErrors
	close(wp.jobCh)
	return wp
}

func (wp *WorkerPool) Process() []error {
	wp.wg.Add(wp.concurrent)
	for i := 0; i < wp.concurrent; i++ {
		go wp.worker()
	}

	go func() {

		for err := range wp.jobResCh {

			if err != nil {
				wp.errors = append(wp.errors, err)
			}

			if len(wp.errors) == wp.maxErrors {
				for i := 0; i < wp.concurrent; i++ {
					wp.jobCancelCh <- true
				}
			}

		}
	}()

	wp.wg.Wait()
	close(wp.jobCancelCh)

	return wp.errors
}
