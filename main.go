package main

import (
	"errors"
	"fmt"
	"otus/api"
	"time"
)

func main() {

	tasks := make([]api.Jobfn, 4)

	tasks[0] = func() error {

		time.Sleep(100 * time.Microsecond)
		return nil
	}

	tasks[1] = func() error {

		fmt.Println("processing task #1")
		time.Sleep(10 * time.Second)
		return errors.New("timeout error33")
	}

	tasks[2] = func() error {

		fmt.Println("processing task #2")
		time.Sleep(10 * time.Second)
		return errors.New("timeout error44")
	}

	tasks[3] = func() error {

		fmt.Println("processing task #3")
		time.Sleep(10 * time.Second)
		return errors.New("timeout error55")
	}

	workerPool := api.NewWorkerPool(tasks, 1, 1)
	jobErrors := workerPool.Process()
	fmt.Println(jobErrors)
}
