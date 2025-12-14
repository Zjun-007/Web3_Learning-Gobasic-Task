/*
基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、
 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    // "demo/constant"
	// "fmt"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)



type Student struct {
    ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    Name  string `gorm:"type:varchar(100);not null" json:"name"`
    Age   int    `gorm:"not null" json:"age"`
    Grade string `gorm:"type:varchar(50);not null" json:"grade"`
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
    dsn := "godb:54862@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	return db
}

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db := ConnectDB()
	err := db.AutoMigrate(&Student{})
	if err != nil {
		panic(err)
	}
	return db
}



func main() {
	db := InitDB()

	// 插入数据
	// db.Create(&Student{ID:1,Name: "张三", Age: 20, Grade: "三年级"})
	// db.Create(&Student{ID:2,Name: "李四", Age: 18, Grade: "二年级"})
	// db.Create(&Student{ID:3,Name: "王五", Age: 12, Grade: "六年级"})

	// 查询数据
	var students []Student
	// db.Find(&students)
	// fmt.Println("所有学生：", students)
	// db.Debug().Where("age > ?", 18).Find(&students)
	// fmt.Println("年龄大于18岁的学生：", students)

	// 更新数据
	// db.Debug().Find(&students).Update("Grade", "四年级").Where("name = ?", "张三")
	// db.Debug().Where("name = ?", "张三").Find(&students).Update("Grade", "二年级")
	// fmt.Println("更新后学生信息：", students[0])

	// 删除数据
	db.Debug().Where("age < ?", 15).Delete(&Student{})
	db.Find(&students)
	fmt.Println("删除后学生信息：", students)

}
