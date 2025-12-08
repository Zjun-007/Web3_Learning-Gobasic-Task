package main

import "fmt"

type Employee struct {
    ID     int
    Name   string
    Salary float64
}

func main() {
    // 创建结构体实例并获取指针
    emp1 := Employee{ID: 1, Name: "Alice", Salary: 50000}
    empPtr := &emp1
    
    // 通过指针访问结构体字段
    fmt.Println("ID:", (*empPtr).ID)  // 显式解引用
    fmt.Println("Name:", empPtr.Name) // 隐式解引用 (Go语法糖)
    
    // 通过指针修改结构体字段
    empPtr.Salary = 55000
    empPtr.Name = "Alice Smith"
    
    fmt.Println("修改后:", emp1)
    
    // 直接创建结构体指针
    emp2 := &Employee{ID: 2, Name: "Bob", Salary: 60000}
    
    // 结构体指针作为函数参数
    giveRaise(emp2, 10) // 加薪10%
    fmt.Println("加薪后:", emp2)
}

// 接收结构体指针参数的函数
func giveRaise(emp *Employee, percent float64) {
    emp.Salary += emp.Salary * percent / 100
}