/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
并将结果映射到 Book 结构体切片中，确保类型安全。
*/

package main	
import (
	"fmt"
	"log"
	"time"		
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)	

// Book 结构体定义
type Book struct {
	ID     int     `db:"id" json:"id"`	
	Title  string  `db:"title" json:"title"`
	Author string  `db:"author" json:"author"`
	Price  float64 `db:"price" json:"price"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at,omitempty"`
}	
// 全局数据库连接
var db *sqlx.DB		
// 初始化数据库连接
func InitDB() error	 {
	dsn := "godb:54862@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=true&loc=Local"
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
// 创建示例表和数据
func CreateSampleData() error {
	// 创建books表
	createTableSQL := `		
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(200) NOT NULL,	
		author VARCHAR(100) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP	
	);`
	_, err := db.Exec(createTableSQL)		
	if err != nil {
		return fmt.Errorf("创建表失败: %w", err)
	}		
	// 清空表数据（可选）
	_, _ = db.Exec("DELETE FROM books")	
	// 插入示例数据
	books := []Book{
		{Title: "Go语言入门", Author: "张三", Price: 45.00},
		{Title: "深入理解Go", Author: "李四", Price: 75.50},
		{Title: "数据库原理", Author: "王五", Price: 60.00},
		{Title: "操作系统概念", Author: "赵六", Price: 80.00},
	}	
	for _, book := range books {
		_, err := db.NamedExec(`INSERT INTO books (title, author, price) VALUES (:title, :author, :price)`, &book)	
		if err != nil {
			return fmt.Errorf("插入数据失败: %w", err)
		}		
	}
	return nil
}
func main() {
	// 初始化数据库
	err := InitDB()	
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	// 创建示例表和数据
	err = CreateSampleData()	
	if err != nil {
		log.Fatalf("创建示例数据失败: %v", err)
	}
	// 查询价格大于50元的书籍
	var expensiveBooks []Book
	query := `SELECT id, title, author, price, created_at FROM books WHERE price > ?`
	err = db.Select(&expensiveBooks, query, 50.0)	
	if err != nil {
		log.Fatalf("查询书籍失败: %v", err)
	}	
	// 输出查询结果	
	fmt.Println("价格大于50元的书籍:")		
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f, CreatedAt: %v\n", book.ID, book.Title, book.Author, book.Price, book.CreatedAt)
	}	
}