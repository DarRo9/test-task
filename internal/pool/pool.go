package pool

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/DarRo9/test-task/internal/stack"
)

var ErrNoWorkers = errors.New("Worker-pool is empty now")

type Worker struct {
	id       int
	cancelCh chan struct{}
}

type Pool struct {
	wg      *sync.WaitGroup
	lastId  atomic.Int64
	workers *stack.Stack[*Worker]
	workCh  chan string
}

func (wp *Pool) AddWork(work string) error {
	if wp.workers.Size() != 0 {
		wp.workCh <- work
		return nil
	}
	return ErrNoWorkers
}

func (wp *Pool) AddWorker() {
	wp.lastId.Add(1)
	worker := &Worker{
		id:       int(wp.lastId.Load()),
		cancelCh: make(chan struct{}),
	}

	wp.workers.Push(worker)
	fmt.Printf("New worker %d added\n", worker.id)

	wp.wg.Add(1)
	go wp.work(worker)
	fmt.Printf("Worker %d started\n", worker.id)
}


func NewPool(numWorkers int) *Pool {
	wp := Pool{
		wg:      &sync.WaitGroup{},
		workers: stack.NewStack[*Worker](),
		workCh:  make(chan string),
	}

	for i := 0; i < numWorkers; i++ {
		wp.AddWorker()
	}

	return &wp
}

func (wp *Pool) work(w *Worker) {
	defer wp.wg.Done()
	for {
		select {
		case <-w.cancelCh:
			return
		case work, ok := <-wp.workCh:
			if !ok {
				fmt.Printf("Worker %d stopped\n", w.id)
				return
			}
			fmt.Printf("Worker %d processing string \"%s\"\n", w.id, work)
		}
	}
}

func (wp *Pool) Stop() {
	close(wp.workCh)
	wp.wg.Wait()
}

func (wp *Pool) RemoveWorker() error {
	worker, err := wp.workers.Pop()
	if err != nil {
		if errors.Is(err, stack.ErrEmptyStack) {
			return ErrNoWorkers
		} else {
			panic(err)
		}
	}

	close(worker.cancelCh)
	fmt.Printf("Worker %d removed\n", worker.id)
	return nil
}
