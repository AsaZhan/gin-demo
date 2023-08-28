package main

import (
	"gin-demo/server/controller/impl"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

var testController = impl.TestController{}

func setupEngine() *gin.Engine {
	//不使用默认中间件
	//r := gin.New()

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	// gin.DisableConsoleColor()

	//强制控制台输出日志使用颜色
	gin.ForceConsoleColor()

	// 记录到文件
	_ = os.Mkdir("log", 0777)
	f, _ := os.Create("log/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//使用默认中间件 Logger Recovery
	r := gin.Default()

	//这玩意影响的是程序启动过程路由绑定的日志，不是调用日志，没啥卵用
	/*gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}*/

	return r
}

func setupRouter() *gin.Engine {

	r := setupEngine()
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

	r.POST("/map", testController.QueryMap)

	r.POST("/cookie", testController.GetAndSetCookie)

	//Schedule
	r.GET("/schedule", testController.Schedule)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080

	//自定义http配置
	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	_ = s.ListenAndServe()

	//r.Run(":8080")
	//log.Fatal(autotls.Run(r,"asa1.dev.com"))
}
