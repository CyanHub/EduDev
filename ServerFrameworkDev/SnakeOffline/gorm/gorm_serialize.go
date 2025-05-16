package gorm

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name       string
	Tags       []string               `gorm:"serializer:json"` // 课程标签
	Syllabus   []string               `gorm:"serializer:json"` // 课程大纲
	Properties map[string]interface{} `gorm:"serializer:json"` // 课程属性
}



func ExampleJSONSerializer() {
    course := &Subject{
        Name: "Go高级编程",
        Tags: []string{"Golang", "后端", "高级"},
        Syllabus: []string{"并发编程", "网络编程", "底层原理"},
        Properties: map[string]interface{}{
            "难度级别": "高级",
            "适合人群": "有Go基础的开发者",
            "预计学时": 48,
        },
    }
    
    // 创建记录
    DB.Create(course)
    
    // 查询记录
    var result Subject
    DB.First(&result, course.ID)
    
    fmt.Printf("课程名称: %s\n", result.Name)
    fmt.Printf("课程标签: %v\n", result.Tags)
    fmt.Printf("课程大纲: %v\n", result.Syllabus)
    fmt.Printf("课程属性: %v\n", result.Properties)
}


type Article struct {
	gorm.Model
	Title string
	Content Content	`gorm:"serializer:gob;type:blob"`
	Likes int
	Views int
}

type Content struct {
	
	Text     string
    Metadata map[string]interface{}

}

func ExampleGobSerializer() {
	article := &Article{
		Title: "Go高级编程",
	}
	var content strings.Builder
    content.WriteString("# 第一章：并发基础\n")
	content.WriteString("## 1.1 并发编程概念\n")
	content.WriteString("Goroutine是Go语言的并发执行单元...\n")
	article.Content = Content{
		Text: content.String(),
		Metadata: map[string]interface{}{
			"author": "张三",
			"published_at": time.Now().Format("2006-01-02 15:04:05"),
			"wordCount":  150,
			"readTime":   "5分钟",
		},
	}
	DB.Create(article)

	var result Article
	DB.First(&result, article.ID)
	fmt.Printf("文章标题: %s\n", result.Title)
	fmt.Printf("文章内容: %s\n", result.Content.Text)
	fmt.Printf("文章元数据: %v\n", result.Content.Metadata)

}
