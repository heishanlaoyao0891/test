package main

// 导入 Gin 包
import "github.com/gin-gonic/gin"

func main() {
	// 1. 创建默认的 Gin 引擎（包含 Logger 和 Recovery 中间件）
	r := gin.Default()

	// 2. 定义 GET 请求路由，路径为 /hello
	r.GET("/hello", func(c *gin.Context) {
		// 3. 返回 JSON 响应
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
			"status":  "success",
		})
	})

	// 4. 启动服务，监听 8080 端口
	// Run() 内部调用了 http.ListenAndServe(":8080", r)
	r.Run(":8082")
}
