 package main

 import (
	 "net/http"
     "github.com/gin-gonic/gin"
 )

func main() {
    // 设置 Gin 模式（可选）
    // gin.SetMode(gin.ReleaseMode)  // 生产模式
    // gin.SetMode(gin.DebugMode)    // 调试模式（默认）
    
    // 创建路由器
    router := gin.Default()
    
    // 1. 根路由 - 返回纯文本
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World!")
    })
    
    // 2. JSON 响应示例
    router.GET("/json", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World",
            "status":  "success",
        })
    })
    
    // 3. HTML 响应示例
    router.GET("/html", func(c *gin.Context) {
        html := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Hello World</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    text-align: center;
                    padding: 50px;
                }
                h1 {
                    color: #333;
                }
            </style>
        </head>
        <body>
            <h1>Hello World</h1>
            <p>Welcome to Gin Framework</p>
        </body>
        </html>
        `
        c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
    })
    
    // 4. 带参数的路由
    router.GET("/hello/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })
    
    // 5. 健康检查端点
    router.GET("/health", func(c *gin.Context) {
        c.String(http.StatusOK, "OK")
    })
    
    // 6. 显示服务器信息
    router.GET("/info", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "server":   "Gin Hello World",
            "version":  "1.0",
            "endpoints": []string{
                "/",
                "/json",
                "/html",
                "/hello/:name",
                "/health",
                "/info",
            },
        })
    })
    
    // 启动服务器
    addr := "127.0.0.1:8080"
    println("🚀 服务器启动中...")
    println("📡 监听地址: http://" + addr)
    println("📋 可用端点:")
    println("  • http://" + addr + "/          - Hello World (纯文本)")
    println("  • http://" + addr + "/json      - Hello World (JSON)")
    println("  • http://" + addr + "/html      - Hello World (HTML)")
    println("  • http://" + addr + "/health    - 健康检查")
    println("  • http://" + addr + "/info      - 服务器信息")
    println("  • http://" + addr + "/hello/:name - 个性化问候")
    
    // 启动服务器
    if err := router.Run(addr); err != nil {
        panic("启动服务器失败: " + err.Error())
    }
}