package main

import (
	"gin-demo/server/controller/impl"
	"github.com/gin-gonic/gin"
)

var testController = impl.TestController{}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	//使用默认中间件 Logger Recovery
	r := gin.Default()

	//不使用默认中间件
	//r := gin.New()

	// Ping test
	r.GET("/ping", testController.Ping)

	// Get user value
	r.GET("/user/:name", testController.GetUserValue)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":    "bar", // user:foo password:bar
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
		"admin":  "adminPw",
	}))
	authorized.POST("/admin", testController.TestAuth)
	authorized.POST("secret", testController.TestSecret)

	// example1
	r.GET("/someJSON", testController.AsciiJSON)

	//example2
	r.GET("/JSONP", testController.CallBack)

	// example3
	r.POST("/login", testController.BindJSON)

	// example4
	r.POST("/form_post", testController.FormData)

	//example5
	r.GET("/json", testController.JSON)
	r.GET("/purejson", testController.PureJSON)

	//example6 query DefaultQuery
	r.POST("/query", testController.Query)

	//secure json
	r.GET("securejson", testController.SecureJSON)

	//XML
	r.GET("/someXML", testController.SomeXML)

	//YAML
	r.GET("/someYAML", testController.SomeYAML)

	//ProtoBuf
	r.GET("/someProtoBuf", testController.SomeProtoBuf)

	//upload
	r.POST("/upload", testController.Upload)

	//upload multi files
	r.POST("/uploadMulti", testController.UploadMulti)

	//getFromReader
	r.POST("/getFromReader", testController.GetFromReader)

	//async
	r.GET("/long_async", testController.Async)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
