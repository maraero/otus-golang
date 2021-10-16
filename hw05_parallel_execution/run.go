package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func handleTask(chTasks chan Task, chErrors chan error, chDone chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-chDone:
			wg.Done()
			return
		case t := <-chTasks:
			if err := t(); err != nil {
				chErrors <- err
			}
		}
	}
}

func Run(tasks []Task, n int, m int) error {
	chTasks := make(chan Task)
	chErrors := make(chan error, len(tasks))
	chDone := make(chan struct{})

	wg := sync.WaitGroup{}

	defer func() {
		close(chDone)
		wg.Wait()
	}()

	wg.Add(n)

	for i := 1; i <= n; i++ {
		go handleTask(chTasks, chErrors, chDone, &wg)
	}

	for _, t := range tasks {
		chTasks <- t

		if m > 0 && len(chErrors) >= m {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
