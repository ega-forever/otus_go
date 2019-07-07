package api

import (
	"sync"
)

type TaskProcessor struct {
	Tasks                   []func() error
	AllowedConcurrentAmount int
	AllowedErrorsAmount     int
	wg                      sync.WaitGroup
	errors                  map[int]error
	errorLockMutex          sync.RWMutex
	sem                     Semaphore
}

func resolver(index int, f func() error, processor *TaskProcessor) {
	defer processor.wg.Done()

	processor.sem.Take(1)

	processor.errorLockMutex.RLock()

	if len(processor.errors) >= processor.AllowedErrorsAmount {
		processor.sem.Release(1)
		processor.errorLockMutex.RUnlock()
		return
	}

	processor.errorLockMutex.RUnlock()

	err := f()
	processor.errorLockMutex.Lock()
	if err != nil {
		processor.errors[index] = err
	}
	processor.errorLockMutex.Unlock()
	processor.sem.Release(1)
}

func (taskProcessor TaskProcessor) Process() map[int]error {

	taskProcessor.errors = make(map[int]error, len(taskProcessor.Tasks))
	taskProcessor.sem = make(Semaphore, taskProcessor.AllowedConcurrentAmount)

	for i, v := range taskProcessor.Tasks {
		taskProcessor.wg.Add(1)
		go resolver(i, v, &taskProcessor)
	}

	taskProcessor.wg.Wait()

	return taskProcessor.errors
}
