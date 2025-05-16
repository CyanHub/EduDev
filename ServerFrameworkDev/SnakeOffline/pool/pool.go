package pool

import (
	"context"
	"errors"
	"sync"
)

type Pool struct {
	tasks  chan func()
	wg     sync.WaitGroup
	size   int
	ctx    context.Context
	cancel context.CancelFunc
}

func NewPool(size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("pool size must be greater than zero")
	}
	ctx, cancel := context.WithCancel(context.Background())
	pool := &Pool{
		tasks:  make(chan func(), size),
		wg:     sync.WaitGroup{},
		size:   size,
		ctx:    ctx,
		cancel: cancel,
	}
	go pool.start()
	return pool, nil
}

func (p *Pool) start() {
	for i := 0; i < p.size; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *Pool) worker() {
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

func (p *Pool) Submit(task func()) {
	select {
	case p.tasks <- task:
	case <-p.ctx.Done():
		return
	}
}

func (p *Pool) Close() {
	p.cancel()
	close(p.tasks)
	p.wg.Wait()
}
