package gorm

import (
	"fmt"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name string
}

type Course struct {
	gorm.Model
	Name      string
	TeacherID *uint
	Teacher   Teacher    `gorm:"foreignKey:TeacherID"`
	Students  []*Student `gorm:"many2many:student_course;"`
}

type Class struct {
	gorm.Model
	Name     string
	Students []*Student
}

type Student struct {
	gorm.Model
	Name    string
	ClassID *uint
	Class   Class     `gorm:"foreignKey:ClassID"`
	Courses []*Course `gorm:"many2many:student_course;"`
}

func AutoMigrate() {
	DB.AutoMigrate(&Teacher{}, &Course{}, &Class{}, &Student{}, &Subject{}, &Article{})
}

func AppendAssociation() {
	// 多对多
	// course := Course{Model: gorm.Model{ID: 16}, Name: "化学"}
	// err := DB.Model(&course).Association("Students").Append([]*Student{
	// 	{Name: "小小"},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// student2 := Student{Model: gorm.Model{ID: 2}, Name: "小红	", ClassID: 1}
	// err = DB.Model(&student2).Association("Courses").Append([]*Course{
	// 	{Model: gorm.Model{ID: 1}, Name: "数学", TeacherID: 1},
	// 	{Model: gorm.Model{ID: 2}, Name: "英语", TeacherID: 2},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// // 不存在会创建
	// teacher := Teacher{Name: "赵六"}
	// err = DB.Create(&teacher).Error
	// if err != nil {
	// 	panic(err)
	// }
	// err = DB.Model(&student).Association("Courses").Append([]*Course{
	// 	{ Name: "物理", TeacherID: teacher.ID},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// 一对一
	// teacher1 := Teacher{Name: "王五"}
	// err = DB.Create(&teacher1).Error
	// if err != nil {
	// 	panic(err)
	// }
	// course := Course{Name: "生物", TeacherID: teacher1.ID}
	// err = DB.Create(&course).Error
	// if err != nil {
	// 	panic(err)
	// }

	// course := Course{Model: gorm.Model{ID: 4}, Name: "物理"}
	// teacher2 := Teacher{Name: "赵六"}
	// err := DB.Model(&course).Association("Teacher").Append(&teacher2)
	// if err != nil {
	// 	panic(err)
	// }

	//  一对多
	// class := Class{Model: gorm.Model{ID: 1}, Name: "一班"}
	// err := DB.Model(&class).Association("Students").Append([]*Student{
	// 	{Model: gorm.Model{ID: 2}, Name: "小红"},
	// 	{Model: gorm.Model{ID: 7}, Name: "小小"},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// class := Class{Model: gorm.Model{ID: 2}, Name: "二班"}
	// err := DB.Model(&class).Association("Students").Append([]*Student{
	// 	{ Name: "丽丽"},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// classId := uint(2)
	// student := Student{Model: gorm.Model{ID: 2}, Name: "小红", ClassID: &classId}
	// err := DB.Model(&student).Association("Class").Append(&Class{ Name: "三班"})
	// if err != nil {
	// 	panic(err)
	// }

}

func QueryAssociation() {
	// course := Course{}
	// DB.First(&course, 1)
	// err := DB.Model(&course).Association("Teacher").Find(&course.Teacher)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("课程名称：%s，教师名称：%s\n", course.Name, course.Teacher.Name)
	// class := Class{}
	// DB.First(&class, 1)
	// err := DB.Model(&class).Association("Students").Find(&class.Students)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, student := range class.Students {
	// 	fmt.Printf("学生名称：%s，班级名称：%s\n", student.Name, class.Name)
	// }
	// student := Student{}
	// DB.First(&student, 1)
	// err := DB.Model(&student).Association("Class").Find(&student.Class)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("学生名称：%s，班级名称：%s\n", student.Name, student.Class.Name)
	// course := Course{}
	// DB.First(&course, 1)
	// err := DB.Model(&course).Association("Students").Find(&course.Students)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, student := range course.Students {
	// 	fmt.Printf("课程名称：%s，学生名称：%s\n", course.Name, student.Name)
	// }

	// student := Student{}
	// DB.First(&student, 1)
	// err := DB.Model(&student).Association("Courses").Find(&student.Courses)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, course := range student.Courses {
	// 	fmt.Printf("学生名称：%s，课程名称：%s\n", student.Name, course.Name)
	// }
}

func ReplaceAssociation() {
	// var course Course
	// DB.First(&course, 4)
	// err := DB.Model(&course).Association("Teacher").Replace(&Teacher{ Name: "lala"})
	// if err != nil {
	// 	panic(err)
	// }

	// student := Student{Model: gorm.Model{ID: 8}, Name: "丽丽"}
	// err := DB.Model(&student).Association("Class").Replace(&Class{Model: gorm.Model{ID: 1}, Name: "一班"})
	// if err != nil {
	// 	panic(err)
	// }

	// student := Student{Model: gorm.Model{ID: 2}, Name: "小红"}
	// err := DB.Model(&student).Association("Courses").Replace([]*Course{
	// 	{Model: gorm.Model{ID: 3}, Name: "生物"},
	// 	{Model: gorm.Model{ID: 4}, Name: "物理"},
	// })
	// if err != nil {
	// 	panic(err)
	// }
}

func CountAssociation() {
	// var course Course
	// DB.First(&course, 16)
	// count := DB.Model(&course).Association("Students").Count()
	// fmt.Printf("课程名称：%s，学生数量：%d\n", course.Name, count)
	var class Class
	DB.First(&class, 1)
	count := DB.Model(&class).Association("Students").Count()
	fmt.Printf("班级名称：%s，学生数量：%d\n", class.Name, count)
}

func DeleteAssociation() {
	// var course Course
	// DB.First(&course, 4)
	// err := DB.Model(&course).Association("Teacher").Find(&course.Teacher)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("课程名称：%s，教师名称：%s\n", course.Name, course.Teacher.Name)
	// err = DB.Model(&course).Association("Teacher").Delete(&course.Teacher)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("课程名称：%s，教师名称：%s\n", course.Name, course.Teacher.Name)
	// var student Student
	// DB.First(&student, 8)
	// err := DB.Model(&student).Association("Class").Find(&student.Class)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("学生名称：%s，班级名称：%s\n", student.Name, student.Class.Name)
	// err = DB.Model(&student).Association("Class").Delete(&student.Class)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("学生名称：%s，班级名称：%s\n", student.Name, student.Class.Name)
	var student Student
	DB.First(&student, 1)
	err := DB.Model(&student).Association("Courses").Find(&student.Courses)
	if err != nil {
		panic(err)
	}
	// for _, course := range student.Courses {
	// 	fmt.Printf("学生名称：%s，课程名称：%s\n", student.Name, course.Name)
	// }
	// fmt.Println("===============================")
	// err = DB.Model(&student).Association("Courses").Delete(student.Courses)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, course := range student.Courses {
	// 	fmt.Printf("学生名称：%s，课程名称：%s\n", student.Name, course.Name)
	// }
}

func AutoAssociation() {
	var course Course
	DB.First(&course, 1)
	student := Student{
		Name: "娜娜",
		Class: Class{
			Name: "五班",
		},
		Courses: []*Course{
			{Name: "历史"},
			{Name: "政治"},
			&course,
		},
	}
	err := DB.Create(&student).Error
	if err != nil {
		panic(err)
	}
}

func PreLoad() {
	var course Course
	DB.First(&course, 1)
	DB.Model(&course).Association("Teacher").Find(&course.Teacher)
	DB.Model(&course).Association("Students").Find(&course.Students)
	for _, student := range course.Students {
		DB.Model(&student).Association("Class").Find(&student.Class)
	}
	// fmt.Println(len(course.Students))
	fmt.Printf("课程名称：%s，教师名称：%s，学生数量：%d\n", course.Name, course.Teacher.Name, len(course.Students))
	for _, student := range course.Students {
		fmt.Printf("学生名称：%s，班级名称：%s\n", student.Name, student.Class.Name)
	}
	// var course Course
	// DB.Preload("Teacher").Preload("Students").Preload("Students.Class").First(&course, 1)
	// fmt.Printf("课程名称：%s，教师名称：%s，学生数量：%d\n", course.Name, course.Teacher.Name, len(course.Students))
	// for _, student := range course.Students {
	// 	fmt.Printf("学生名称：%s，班级名称：%s\n", student.Name, student.Class.Name)
	// }

}
