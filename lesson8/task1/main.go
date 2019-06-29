package main

import (
	"errors"
	"fmt"
	"otus/lesson8/task1/api"
	"time"
)

func main() {

	tasks := make([]func() error, 4)

	tasks[0] = func() error {

		time.Sleep(100 * time.Microsecond)
		return nil
	}

	tasks[1] = func() error {

		time.Sleep(10 * time.Second)
		return errors.New("timeout error33")
	}

	tasks[2] = func() error {

		time.Sleep(10 * time.Second)
		return errors.New("timeout error44")
	}

	tasks[3] = func() error {

		time.Sleep(10 * time.Second)
		return errors.New("timeout error55")
	}

	pr := api.TaskProcessor{Tasks: tasks, AllowedErrorsAmount: 1, AllowedConcurrentAmount: 2}
	prResults := pr.Process()
	fmt.Println(prResults)

}
