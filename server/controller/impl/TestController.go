package impl

import (
	"fmt"
	"gin-demo/server/models/login"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"os"
	"time"
)

//var db = make(map[string]string)
var db = map[string]interface{}{
	"foo":    login.User{Value: "foo", Email: "foo@bar.com", Phone: "123433"},
	"austin": login.User{Value: "austin", Email: "austin@example.com", Phone: "666"},
	"lena":   login.User{Value: "lena", Email: "lena@guapa.com", Phone: "523443"},
}

type TestController struct{}

//Ping test
func (controller *TestController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Get user value
func (controller *TestController) GetUserValue(c *gin.Context) {
	db["张三"] = "法外狂徒"
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

//BasicAuth中间件1
func (controller *TestController) TestAuth(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	// Parse JSON
	request := login.User{}
	err := c.Bind(&request)

	if err == nil {
		db[user] = request.Value
		response := login.Response{Message: "success", Status: "ok"}
		fmt.Printf("%+v\n", db)
		c.JSON(http.StatusOK, response)
	}
}

//BasicAuth中间件2
func (controller *TestController) TestSecret(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	secret, ok := db[user]
	response := login.Response{}
	if ok {
		response = login.Response{
			Message: "Success",
			Status:  "ok",
			Secret:  secret,
		}

	} else {
		response = login.Response{
			Message: "Success",
			Status:  "ok",
			Secret:  "no secret",
		}
	}
	c.JSON(http.StatusOK, response)
}

func (controller *TestController) AsciiJSON(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

func (controller *TestController) CallBack(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	// /JSONP?callback=x
	// 将输出：x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}

func (controller *TestController) BindJSON(c *gin.Context) {
	var form login.LoginForm
	// 你可以使用显式绑定声明绑定 multipart form：
	//c.ShouldBindWith(&form, binding.Form)

	//⬇ 绑定请求参数为JSON类型或结构体，请求参数位于body️
	//if c.ShouldBind(&form) == nil {
	if c.ShouldBindWith(&form, binding.JSON) == nil {
		response := login.Response{Message: "success", Status: "ok"}
		if form.User == "admin" && form.Password == "adminPw" {
			response.Status = "you are logged in"
			c.JSON(http.StatusOK, response)
		} else {
			response.Status = "unauthorized"
			c.JSON(http.StatusUnauthorized, response)
		}
	}
}

//form-data表单传参  DefaultPostForm  PostForm
//参数位于body ContentType 为form-data
func (controller *TestController) FormData(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

//json purejson
// curl命令行可显式区别，postman会自动转码
func (controller *TestController) JSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

func (controller *TestController) PureJSON(c *gin.Context) {
	c.PureJSON(http.StatusOK, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

//Content-Type: application/x-www-form-urlencoded
// Query和DefaultQuery参数位于请求路径param
//ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据
func (controller *TestController) Query(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.DefaultPostForm("name", "张三")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
}

func (controller *TestController) SecureJSON(c *gin.Context) {
	name := []string{"hello", "golang", "你好", "123"}
	c.SecureJSON(http.StatusOK, name)
}

//响应xml格式
func (controller *TestController) SomeXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"Message": "success", "Type": "XML"})
}

//响应yaml格式
func (controller *TestController) SomeYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"Message": "success", "Type": "YAML"})
}

func (controller *TestController) SomeProtoBuf(c *gin.Context) {
	label := "label"
	data := &login.ProtoBufExample{
		Label: label,
		Resp:  []int64{int64(1), int64(2)},
	}

	c.ProtoBuf(http.StatusOK, data)
}

//单一上传
func (controller *TestController) Upload(c *gin.Context) {
	//获取工作目录根目录路径
	wd, _ := os.Getwd()
	//从参数中接收文件
	file, _ := c.FormFile("file")
	filename := file.Filename

	//目标路径
	des := wd + "/resources/" + filename
	//保存文件
	_ = c.SaveUploadedFile(file, des)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", filename))
}

//批量上传
func (controller *TestController) UploadMulti(c *gin.Context) {
	//获取工作目录根目录路径
	wd, _ := os.Getwd()
	//从参数中接收文件
	MultiPartFile, _ := c.MultipartForm()
	files := MultiPartFile.File["files"]
	for _, file := range files {
		filename := file.Filename
		//目标路径
		des := wd + "/resources/" + filename
		//保存文件
		_ = c.SaveUploadedFile(file, des)
	}
	c.String(http.StatusOK, fmt.Sprintf("'%d' file uploaded!", len(files)))
}

//从Reader中取数据
func (controller *TestController) GetFromReader(c *gin.Context) {
	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	if err != nil || http.StatusOK != response.StatusCode {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

//模拟异步响应
func (controller *TestController) Async(c *gin.Context) {
	//cp := c.Copy()
	chn := make(chan string)
	go mockAsync(chn, c)
	log.Println(<-chn)
	c.JSON(http.StatusOK, login.Response{
		Message: "Success",
		Status:  "OK",
		Secret:  nil,
	})
}

func mockAsync(chn chan<- string, c *gin.Context) {
	var inf = "Done! in path " + c.Request.URL.Path
	time.Sleep(5 * time.Second)
	chn <- inf
}

//映射查询字符串或表单参数,映射为map
func (controller *TestController) QueryMap(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	log.Printf("ids: %v; names: %v\n", ids, names)
	c.JSON(http.StatusOK, login.Response{
		Message: "Success",
		Status:  "OK",
		Secret:  nil,
	})
}

func (controller TestController) GetAndSetCookie(c *gin.Context) {
	cookie, err := c.Cookie("DataSet")
	if err != nil {
		log.Println("Set Cookie...")
		c.SetCookie("DataSet", time.Now().String(), 300, "/", "localhost", false, true)
	}
	c.JSON(http.StatusOK, login.Response{
		Message: "Success",
		Status:  "OK",
		Secret:  cookie,
	})
}
