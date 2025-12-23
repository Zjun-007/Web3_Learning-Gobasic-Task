个人博客系统后端
基于 Go + Gin + GORM 开发的个人博客系统后端，支持用户认证、文章管理和评论功能。

功能特性
✅ 用户注册与登录（JWT认证）

✅ 文章CRUD操作

✅ 文章权限控制（只有作者可修改/删除）

✅ 评论功能

✅ 统一的错误处理和日志记录

✅ MySQL数据库支持

技术栈
Go 1.21+

Gin - Web框架

GORM - ORM库

MySQL - 数据库

JWT - 身份认证

bcrypt - 密码加密

项目结构
text
blog/
├── cmd/
│   └── server/
│       └── main.go          # 应用入口
├── internal/
│   ├── config/
│   │   └── config.go        # 配置管理
│   ├── database/
│   │   └── database.go      # 数据库连接
│   ├── models/
│   │   ├── user.go          # 用户模型
│   │   ├── post.go          # 文章模型
│   │   └── comment.go       # 评论模型
│   ├── controllers/
│   │   ├── auth.go          # 认证处理器
│   │   ├── post.go          # 文章处理器
│   │   └── comment.go       # 评论处理器
│   ├── middleware/
│   │   ├── auth.go          # 认证中间件
│   │   └── logger.go        # 日志中间件
│   ├── utils/
│   │   ├── jwt.go           # JWT工具
│   │   └── password.go      # 密码工具
│   └── pkg/
│       └── response/
│           └── response.go   # 统一响应格式
├── pkg/
│   └── logger/
│       └── logger.go        # 日志封装
├── storage/
│   └── logs/
│       └── app.log          # 日志文件
├── .env.example            # 环境变量示例
├── .gitignore
├── go.mod
├── go.sum
└── README.md
快速开始
1. 环境要求
Go 1.21+

MySQL 8.0+

Git

2. 克隆项目
bash
git clone <repository-url>
cd blog-backend
3. 配置环境变量
复制环境变量示例文件：

bash
cp .env.example .env
编辑 .env 文件，配置数据库连接等信息：

env
# 服务器配置
SERVER_HOST=localhost
SERVER_PORT=8080

# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=blog_db

# JWT配置
JWT_SECRET=your-jwt-secret-key-change-this
JWT_EXPIRE=24h
4. 初始化数据库
登录 MySQL，创建数据库：

sql
CREATE DATABASE IF NOT EXISTS blog_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
5. 安装依赖
bash
go mod tidy
6. 运行项目
bash
# 开发模式（使用热重载）
go run cmd/server/main.go

# 或者直接运行
go build -o blog-backend cmd/server/main.go
./blog-backend
服务器将在 http://localhost:8080 启动。

API 接口文档
认证相关
用户注册
URL: /api/v1/auth/register

Method: POST

认证: 不需要

请求体:

json
{
    "username": "john_doe",
    "password": "password123",
    "email": "john@example.com"
}
用户登录
URL: /api/v1/auth/login

Method: POST

认证: 不需要

请求体:

json
{
    "username": "john_doe",
    "password": "password123"
}
响应:

json
{
    "code": 200,
    "message": "登录成功",
    "data": {
        "token": "jwt-token-here",
        "user": {
            "id": 1,
            "username": "john_doe",
            "email": "john@example.com"
        }
    }
}
文章管理
获取所有文章
URL: /api/v1/posts

Method: GET

认证: 不需要

查询参数:

page (可选): 页码，默认1

limit (可选): 每页数量，默认10

获取单篇文章
URL: /api/v1/posts/:id

Method: GET

认证: 不需要

创建文章
URL: /api/v1/posts

Method: POST

认证: 需要（Bearer Token）

请求体:

json
{
    "title": "文章标题",
    "content": "文章内容..."
}
更新文章
URL: /api/v1/posts/:id

Method: PUT

认证: 需要（Bearer Token，只能更新自己的文章）

请求体:

json
{
    "title": "更新后的标题",
    "content": "更新后的内容..."
}
删除文章
URL: /api/v1/posts/:id

Method: DELETE

认证: 需要（Bearer Token，只能删除自己的文章）

评论管理
获取文章评论
URL: /api/v1/posts/:postId/comments

Method: GET

认证: 不需要

创建评论
URL: /api/v1/posts/:postId/comments

Method: POST

认证: 需要（Bearer Token）

请求体:

json
{
    "content": "评论内容"
}
测试用例
使用 Postman 测试
导入 Postman 集合

创建新的 Collection

添加环境变量 base_url = http://localhost:8080/api/v1

添加环境变量 token（登录后更新）

测试流程

text
1. 注册用户
2. 登录获取token
3. 创建文章
4. 获取文章列表
5. 更新文章
6. 添加评论
7. 获取评论
8. 删除文章
示例测试命令
bash
# 测试注册
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"test123","email":"test@example.com"}'

# 测试登录
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"test123"}'

# 使用token创建文章
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title":"测试文章","content":"这是测试文章的内容"}'

# 获取文章列表
curl -X GET http://localhost:8080/api/v1/posts
错误处理
系统使用统一的错误响应格式：

json
{
    "code": 400,
    "message": "错误描述",
    "data": null
}
常见HTTP状态码
200: 成功

201: 创建成功

400: 请求参数错误

401: 未授权

403: 禁止访问（权限不足）

404: 资源不存在

500: 服务器内部错误

日志记录
日志文件位于 storage/logs/app.log，包含：

请求信息（IP、方法、路径、状态码）

错误堆栈

数据库查询日志（开发环境）

部署
Docker 部署
创建 Dockerfile:

dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

CMD ["./main"]
构建并运行：

bash
docker build -t blog-backend .
docker run -p 8080:8080 blog-backend
生产环境建议
修改 .env 中的配置：

使用强JWT密钥

配置生产数据库

设置适当的日志级别

使用反向代理（Nginx）

配置HTTPS

设置数据库备份

监控和告警

开发说明
添加新功能
在 models/ 中添加新的模型

在 handlers/ 中添加处理器

在 routers/ 中注册路由

添加相关测试

数据库迁移
系统使用 GORM AutoMigrate，首次运行会自动创建表结构。

手动查看数据库状态：

bash
# 查看表结构
mysql -u root -p blog_db -e "SHOW TABLES;"

# 查看用户表结构
mysql -u root -p blog_db -e "DESCRIBE users;"
故障排除
常见问题
数据库连接失败

检查数据库服务是否运行

检查 .env 中的数据库配置

检查防火墙设置

JWT认证失败

检查Token是否过期

检查Authorization头格式

验证JWT_SECRET配置

权限错误

确保用户已登录

检查是否是资源所有者

查看日志
bash
# 查看实时日志
tail -f storage/logs/app.log

# 查看错误日志
grep "ERROR" storage/logs/app.log
贡献指南
Fork 项目

创建功能分支

提交更改

推送到分支

创建 Pull Request

许可证
MIT License