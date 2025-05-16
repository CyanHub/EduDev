package crontab

import (
	"fmt"
	"time"
)

func PrintTime() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()


	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("定时任务执行：", time.Now())
			case <-quit:
				return
			}
		}
	}()

	time.Sleep(10 * time.Second)
	close(quit)

}
