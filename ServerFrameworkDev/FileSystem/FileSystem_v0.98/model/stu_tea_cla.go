package model

import "gorm.io/gorm"

// Teacher 表示教师信息的结构体。
type Teacher struct {
	gorm.Model
	Name string
}

// Course 表示课程信息的结构体，关联教师和学生。
type Course struct {
	gorm.Model
	Name      string
	TeacherID uint
	Teacher   Teacher   `gorm:"foreignKey:TeacherID"`
	Students  []Student `gorm:"many2many:student_course;"`
}

// Class 表示班级信息的结构体，包含多个学生。
type Class struct {
	gorm.Model
	Name     string
	Students []Student
}

// Student 表示学生信息的结构体，关联班级和课程。
type Student struct {
	gorm.Model
	Name    string
	ClassID uint
	Class   Class
	Courses []Course `gorm:"many2many:student_course;"`
}

// Subject 表示学科信息的结构体，包含学科名称、标签、大纲和属性。
// type Subject struct {
// 	gorm.Model
// 	Name       string
// 	Tags       []string               `gorm:"serializer:json"` // 课程标签
// 	Syllabus   []string               `gorm:"serializer:text"` // 课程大纲
// 	Properties map[string]interface{} `gorm:"serializer:json"` // 课程属性
// }

// Article 表示文章信息的结构体，包含标题、内容、浏览次数和点赞次数。
type Article struct {
	gorm.Model
	Title   string  // 标题
	Content Content `gorm:"serializer:gob;type:blob"` // 内容
	Views   int     // 浏览次数
	Likes   int     // 点赞次数
}

// Content 表示文章内容的结构体，包含文本内容和元数据。
type Content struct {
	Text     string                 // 文章内容
	MetaData map[string]interface{} // 元数据
}
