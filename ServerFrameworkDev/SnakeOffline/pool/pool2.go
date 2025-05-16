package pool

import (
	"context"
	"errors"
	"sync"
)

type Pool2 struct {
	tasks  chan func()
	wg     sync.WaitGroup
	size   int
	ctx    context.Context
	cancel context.CancelFunc
}

func NewPool2(size int) (*Pool2, error) {
	if size <= 0 {
		return nil, errors.New("pool size must be greater than zero")
	}
	ctx, cancel := context.WithCancel(context.Background())
	pool := &Pool2{
		tasks:  make(chan func(), size),
		wg:     sync.WaitGroup{},
		size:   size,
		ctx:    ctx,
		cancel: cancel,
	}
	go pool.start()
	return pool, nil
}

func (p *Pool2) start() {
	for i := 0; i < p.size; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *Pool2) worker() {
	defer p.wg.Done()
	for {
		select {
		case task := <-p.tasks:
			task()
		case <-p.ctx.Done():
			return
		}
	}
}

func (p *Pool2) Submit(task func()) {
	p.tasks <- task
}

func (p *Pool2) Wait() {
	p.wg.Wait()
}

func (p *Pool2) Close() {
	p.cancel()
	close(p.tasks)
	p.Wait()
}
