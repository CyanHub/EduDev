package main

import (
	"fmt"
	"log"
	"ServerFramework/pool"
	"sync"
	"time"
)

func main() {
	// viper.ReadFlag()
	// viper.ReadConfigWithHotReload()
	// select {}
	// zap.ZapWithFile()
	// gorm.AutoMigrate()
	// gorm.ReplaceAssociation()
	// crontab.PrintTime()
	// 创建一个大小为 5 的协程池
	p, err := pool.NewPool(5)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer p.Close()

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		taskID := i
		p.Submit(func() {
			defer wg.Done()
			fmt.Printf("Task %d started\n", taskID)
			time.Sleep(time.Second)
			fmt.Printf("Task %d finished\n", taskID)
		})
	}
	wg.Wait()
	fmt.Println("All tasks completed")
}
