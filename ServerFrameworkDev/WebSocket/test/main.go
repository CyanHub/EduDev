package test

// 测试用 main组件

import (
	"fmt"
	"os"
	"strings"
	"time"
	
	"ServerFramework/global"
	"ServerFramework/model"

	"github.com/robfig/cron/v3"
)

func ExampleJSONSerializer() {
	course := &model.Subject{
		Name:     "Go高级编程1000",
		Tags:     []string{"Golang", "后端", "高级"},
		Syllabus: []string{"并发编程", "网络编程", "底层原理"},
		Properties: map[string]interface{}{
			"难度级别": "高级",
			"适合人群": "有Go基础的开发者",
			"预计学时": 48,
		},
	}

	// 创建记录
	global.DB.Create(course)
	// 查询记录
	var result model.Subject

	global.DB.First(&result, course.ID)

	fmt.Printf("课程名称: %s\n", result.Name)
	fmt.Printf("课程标签: %v\n", result.Tags[0])
	fmt.Printf("课程大纲: %v\n", result.Syllabus)
	fmt.Printf("课程属性: %v\n", result.Properties)
}

func ExampleGobSerializer() {
	article := &model.Article{
		Title: "Go高级编程",
	}
	var content strings.Builder
	content.WriteString("# 第一章：并发基础\n")

	content.WriteString(" # 1.1 并发编程概念\n")
	content.WriteString("Goroutine是Go语言的并发执行单元 .\n")
	article.Content = model.Content{
		Text: content.String(),
		MetaData: map[string]interface{}{
			"author":       "张三",
			"published_at": time.Now().Format("2006-01-02 15:04:05"),
			"wordCount":    len(content.String()),
			"readTime":     "5分钟",
		},
	}
	global.DB.Create(article)

	var res model.Article
	global.DB.First(&res, article.ID)

	file, err := os.OpenFile("./article.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	n, err := file.WriteString(res.Content.Text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入了%d个字节", n)

}

func PrintTime() /*error*/ {
	fmt.Println("当前时间是：", time.Now().Format("2006-01-02 15:04:05"))
	// return nil
}

func CronUse() {
	c := cron.New(cron.WithSeconds())
	c.Start()

	_, err := c.AddFunc("*/2 * * * * *", func() {
		fmt.Println("当前时间是：", time.Now().Format("2006-01-02 15:04:05"))
		//PrintTime()
	})
	if err != nil {
		panic(fmt.Sprintf("添加定时任务失败: %v", err))
	}
	time.Sleep(2 * time.Second)

	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("定时任务退出")
				break
			}
		}
	}()
	time.Sleep(10 * time.Second)
	quit <- struct{}{}
	time.Sleep(5 * time.Second)
}

func TickerUse() {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("当前时间是：", time.Now().Format("2006-01-02 15:04:05"))
				break
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(10 * time.Second)
	quit <- struct{}{}
}
