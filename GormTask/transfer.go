/*
事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， 
amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
如果余额不足，则回滚事务。 
*/
package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Account 账户表
type Account struct {
	ID      uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Balance float64 `gorm:"not null;default:0;check:balance>=0" json:"balance"`
}

// Transaction 交易记录表
type Transaction struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FromAccountID uint      `gorm:"not null;index" json:"from_account_id"`
	ToAccountID   uint      `gorm:"not null;index" json:"to_account_id"`
	Amount        float64   `gorm:"not null;check:amount>0" json:"amount"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	dsn := "godb:54862@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 启用预编译语句
	})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	return db
}

// InitDB 初始化数据库和表
func InitDB() *gorm.DB {
	// db := ConnectDB()
	
	// // 设置连接池
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatal("获取数据库连接失败:", err)
	// }
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	
	// // 自动迁移表结构
	// err = db.AutoMigrate(&Account{}, &Transaction{})
	// if err != nil {
	// 	log.Fatal("数据库迁移失败:", err)
	// }
	
	// log.Println("数据库初始化成功")
	// return db
	db := ConnectDB()
	err := db.AutoMigrate(&Account{}, &Transaction{})
	if err != nil {
		panic(err)
	}
	return db
}

// CreateSampleAccounts 创建示例账户
func CreateSampleAccounts(db *gorm.DB) {
	// 创建账户A，初始余额500元
	accountA := Account{ID: 1, Balance: 500.0}
	// 创建账户B，初始余额200元
	accountB := Account{ID: 2, Balance: 200.0}
	
	// 使用FirstOrCreate避免重复创建
	db.FirstOrCreate(&accountA, Account{ID: 1})
	db.FirstOrCreate(&accountB, Account{ID: 2})
	
	log.Printf("账户A(ID: %d) 余额: %.2f\n", accountA.ID, accountA.Balance)
	log.Printf("账户B(ID: %d) 余额: %.2f\n", accountB.ID, accountB.Balance)
}

// TransferMoney 转账函数
func TransferMoney(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	// 参数校验
	if fromAccountID == toAccountID {
		return fmt.Errorf("转出账户和转入账户不能相同")
	}
	if amount <= 0 {
		return fmt.Errorf("转账金额必须大于0")
	}
	
	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("事务启动失败: %w", tx.Error)
	}
	
	// 声明在函数结束时检查是否需要回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("事务因panic回滚: %v", r)
		}
	}()
	
	// 1. 使用悲观锁查询转出账户（防止并发问题）
	var fromAccount Account
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&fromAccount, fromAccountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %w", err)
	}
	
	// 2. 检查余额是否足够
	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足，当前余额: %.2f", fromAccount.Balance)
	}
	
	// 3. 使用悲观锁查询转入账户
	var toAccount Account
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&toAccount, toAccountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转入账户失败: %w", err)
	}
	
	// 4. 更新转出账户余额
	if err := tx.Model(&fromAccount).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新转出账户余额失败: %w", err)
	}
	
	// 5. 更新转入账户余额
	if err := tx.Model(&toAccount).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新转入账户余额失败: %w", err)
	}
	
	// 6. 在transactions表中记录转账信息
	transaction := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录转账信息失败: %w", err)
	}
	
	// 7. 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("事务提交失败: %w", err)
	}
	
	log.Printf("转账成功: 从账户%d向账户%d转账%.2f元", fromAccountID, toAccountID, amount)
	log.Printf("交易记录ID: %d", transaction.ID)
	return nil
}

// GetAccountInfo 获取账户信息
func GetAccountInfo(db *gorm.DB, accountID uint) (*Account, error) {
	var account Account
	if err := db.First(&account, accountID).Error; err != nil {
		return nil, fmt.Errorf("账户不存在: %w", err)
	}
	return &account, nil
}

func main() {
	// 1. 初始化数据库
	db := InitDB()
	
	// 2. 创建示例账户
	CreateSampleAccounts(db)
	
	// 3. 显示转账前账户余额
	fmt.Println("=== 转账前账户余额 ===")
	accountA, _ := GetAccountInfo(db, 1)
	accountB, _ := GetAccountInfo(db, 2)
	fmt.Printf("账户A(ID: %d) 余额: %.2f元\n", accountA.ID, accountA.Balance)
	fmt.Printf("账户B(ID: %d) 余额: %.2f元\n", accountB.ID, accountB.Balance)
	
	// 4. 执行转账操作（从账户A向账户B转账100元）
	fmt.Println("\n=== 执行转账操作 ===")
	fromAccountID := uint(1) // 账户A
	toAccountID := uint(2)   // 账户B
	amount := 100.0
	
	fmt.Printf("尝试转账: 从账户%d向账户%d转账%.2f元\n", fromAccountID, toAccountID, amount)
	
	err := TransferMoney(db, fromAccountID, toAccountID, amount)
	if err != nil {
		fmt.Printf("转账失败: %v\n", err)
	} else {
		fmt.Println("转账成功!")
	}
	
	// 5. 显示转账后账户余额
	fmt.Println("\n=== 转账后账户余额 ===")
	accountA, _ = GetAccountInfo(db, 1)
	accountB, _ = GetAccountInfo(db, 2)
	fmt.Printf("账户A(ID: %d) 余额: %.2f元\n", accountA.ID, accountA.Balance)
	fmt.Printf("账户B(ID: %d) 余额: %.2f元\n", accountB.ID, accountB.Balance)
	
	// 6. 测试余额不足的情况
	fmt.Println("\n=== 测试余额不足的情况 ===")
	err = TransferMoney(db, 1, 2, 500.0) // 账户A余额不足以转账500元
	if err != nil {
		fmt.Printf("预期中的转账失败: %v\n", err)
	}

	fmt.Println("\n=== 再次查询账户余额 ===")
	accountA, _ = GetAccountInfo(db, 1)
	accountB, _ = GetAccountInfo(db, 2)
	fmt.Printf("账户A(ID: %d) 余额: %.2f元\n", accountA.ID, accountA.Balance)
	fmt.Printf("账户B(ID: %d) 余额: %.2f元\n", accountB.ID, accountB.Balance)

	
	// 7. 显示交易记录
	fmt.Println("\n=== 最近5笔交易记录 ===")
	var transactions []Transaction
	db.Order("created_at DESC").Limit(5).Find(&transactions)
	
	for _, tx := range transactions {
		fmt.Printf("交易ID: %d | 从账户%d向账户%d转账%.2f元 | 时间: %s\n",
			tx.ID, tx.FromAccountID, tx.ToAccountID, tx.Amount,
			tx.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}