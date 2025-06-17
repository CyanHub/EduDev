package utils

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// 定义协程池结构体
type Pool struct {
	Tasks  chan func() int    // 任务管道
	Size   int                // 任务数量
	Wg     sync.WaitGroup     // 并发控制组件
	Ctx    context.Context    // 上下文
	Cancel context.CancelFunc // 取消信号
}

// NewPool 创建一个新的协程池
// size: 协程池的大小
// 返回值: *Pool 协程池实例, error 错误信息
func NewPool(size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("协程池中的协程数量不能小于1")
	}

	// 创建一个带有取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 初始化协程池
	pool := &Pool{
		Tasks:  make(chan func() int, size+2), // 创建任务通道，容量为size+2
		Wg:     sync.WaitGroup{},              // 初始化WaitGroup
		Size:   size,                          // 设置协程池大小
		Ctx:    ctx,                           // 设置上下文
		Cancel: cancel,                        // 设置取消函数
	}

	// 启动协程池中的工作协程
	go pool.Start()
	return pool, nil
}

// Start 启动工作Goroutine
func (p *Pool) Start() {
	// 循环启动指定数量的工作协程
	for i := 1; i < p.Size; i++ {
		p.Wg.Add(1)  // 增加WaitGroup计数
		go p.Work(i) // 启动工作协程
	}
}

// Work 工作协程
// i: 协程编号
func (p *Pool) Work(i int) {
	defer p.Wg.Done() // 在函数结束时调用Done，通知WaitGroup当前协程已完成

	for {
		select {
		case task := <-p.Tasks: // 从任务通道中接收任务
			taskID := task()                       // 执行任务并获取任务ID
			fmt.Printf("协程%d执行了任务%d\n", i, taskID) // 打印任务执行信息
		case <-p.Ctx.Done(): // 接收到上下文取消信号
			fmt.Println("接收到协程取消信号，协程", i, "退出") // 打印取消信息
			return                               // 结束协程
		}
	}
}

// Submit 提交任务到协程池
// task: 需要执行的任务，任务是一个返回整数的函数
// 返回值: error 错误信息，如果协程池已关闭则返回错误
func (p *Pool) Submit(task func() int) error {
	select {
	case <-p.Ctx.Done(): // 检查上下文是否已取消
		return fmt.Errorf("协程池已关闭，无法提交数据")
	case p.Tasks <- task: // 将任务发送到任务通道
		fmt.Println("提交任务成功")
	}
	return nil
}

// 关闭协程池
func (p *Pool) Close() {
	p.Cancel()
	close(p.Tasks)
	p.Wg.Wait()
}
