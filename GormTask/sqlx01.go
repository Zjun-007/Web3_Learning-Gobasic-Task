/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，
包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

package main

import (
    // "context"
    // "database/sql"
    "fmt"
    "log"
    "time"

    _"github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

// Employee 结构体定义
type Employee struct {
    ID         int     `db:"id" json:"id"`
    Name       string  `db:"name" json:"name"`
    Department string  `db:"department" json:"department"`
    Salary     float64 `db:"salary" json:"salary"`
    CreatedAt  *time.Time `db:"created_at" json:"created_at,omitempty"`
}

// 全局数据库连接
var db *sqlx.DB

// 初始化数据库连接
func InitDB() error {
    dsn := "godb:54862@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=true&loc=Local"
    
    var err error
    db, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        return fmt.Errorf("数据库连接失败: %w", err)
    }
    
    // 配置连接池
    db.SetMaxOpenConns(50)
    db.SetMaxIdleConns(10)
    db.SetConnMaxLifetime(time.Hour)
    
    return nil
}

// 创建示例表和数据
func CreateSampleData() error {
    // 创建employees表
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS employees (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        department VARCHAR(50) NOT NULL,
        salary DECIMAL(10, 2) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
    
    _, err := db.Exec(createTableSQL)
    if err != nil {
        return fmt.Errorf("创建表失败: %w", err)
    }
    
    // 清空表数据（可选）
    // _, _ = db.Exec("DELETE FROM employees")
    
    // 插入示例数据
    employees := []Employee{
        {Name: "张三", Department: "技术部", Salary: 15000.00},
        {Name: "李四", Department: "技术部", Salary: 18000.00},
        {Name: "王五", Department: "市场部", Salary: 12000.00},
        {Name: "赵六", Department: "技术部", Salary: 20000.00},
        {Name: "钱七", Department: "人事部", Salary: 10000.00},
        {Name: "孙八", Department: "技术部", Salary: 22000.00},
    }
    
    for _, emp := range employees {
        query := `INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`
        _, err := db.Exec(query, emp.Name, emp.Department, emp.Salary)
        if err != nil {
            return fmt.Errorf("插入数据失败: %w", err)
        }
    }
    
    log.Println("示例数据创建完成")
    return nil
}

func main() {
	// 初始化数据库
	err := InitDB()			
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}				
	defer db.Close()

	// 创建示例数据
	err = CreateSampleData()
	if err != nil {
		log.Fatalf("创建示例数据失败: %v", err)
	}				
	// 查询技术部所有员工
	var techEmployees []Employee
	query1 := `SELECT id, name, department, salary, created_at FROM employees WHERE department = ?`
	err = db.Select(&techEmployees, query1, "技术部")		
	if err != nil {
		log.Fatalf("查询技术部员工失败: %v", err)
	}		
	fmt.Println("技术部员工列表:")
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f, CreatedAt: %v\n", emp.ID, emp.Name, emp.Department, emp.Salary, emp.CreatedAt)
	}
	// 查询工资最高的员工		
	var topEmployee Employee		
	query2 := `SELECT id, name, department, salary, created_at FROM employees ORDER BY salary DESC LIMIT 1`
	err = db.Get(&topEmployee, query2)
	if err != nil {
		log.Fatalf("查询工资最高的员工失败: %v", err)
	}				
	fmt.Println("\n工资最高的员工:")
	fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f, CreatedAt: %v\n", topEmployee.ID, topEmployee.Name, topEmployee.Department, topEmployee.Salary, topEmployee.CreatedAt)
}	