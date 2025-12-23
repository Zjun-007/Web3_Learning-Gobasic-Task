package main

import (
    "blog/config"
    "blog/database"
    // "blog/middleware"
    "blog/routes"
    "log"
	"os"
	// "github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

    // "github.com/gin-gonic/gin"
)

func main() {
    // 加载配置
    config.LoadConfig()
    
    // 连接数据库
    database.InitDB()
    
    // 设置Gin模式
    // if config.AppConfig == "production" {
    //     gin.SetMode(gin.ReleaseMode)
    // } else {
    //     gin.SetMode(gin.DebugMode)
    // }
    
	// 设置路由
	r := routes.SetupRoutes()

    // // 创建Gin实例
    // router := gin.Default()
    
    // 设置模板函数
    // router.SetFuncMap(routes.TemplateFuncMap())
    
    // // 加载模板
    // router.LoadHTMLGlob("templates/**/*.html")
	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	
	logrus.WithField("port", port).Info("Server starting")
	
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
    
    
}