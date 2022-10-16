// Package hw05parallelexecution -- HW05 Otus Go.
package hw05parallelexecution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

// ErrErrorsLimitExceeded -- limit exceeded.
var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

// Task -- simple task type.
type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, n int, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	errAcc := int32(1)
	errMax := int32(m)
	ch := make(chan Task, len(tasks))
	wg := &sync.WaitGroup{}

	// Workers.
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			acc := int32(1)
			for t := range ch {
				err := t()
				if err != nil {
					acc = atomic.AddInt32(&errAcc, 1)
				}
				if acc > errMax {
					return
				}
			}
		}()
	}

	// Tasks.
	for _, t := range tasks {
		ch <- t
	}

	close(ch)
	wg.Wait()

	if errAcc > errMax {
		return ErrErrorsLimitExceeded
	}
	return nil
}
