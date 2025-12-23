package database

import (
    "fmt"
    // "log"
	"blog/config"
    "blog/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    // "gorm.io/gorm/logger"
)


var DB *gorm.DB

func ConnectDB() *gorm.DB {
    cfg := config.AppConfig
    
    // 构建MySQL连接字符串
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
    )

    DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	return DB
}

// func ConnectDB() *gorm.DB {
//     dsn := "godb:54862@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
//     db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect database")
//     }
// 	return db
// }

// func GetDB() *gorm.DB {
// 	db := ConnectDB()
// 	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
//     if err != nil {
//         return fmt.Errorf("failed to migrate database: %w", err)
//     }
//     return db
// }


func InitDB() *gorm.DB {
	DB := ConnectDB()
	err := DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		panic(err)
	}
	return DB
}