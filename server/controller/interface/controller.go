package interfaces

import "github.com/gin-gonic/gin"

type TestController interface {
	Ping(c *gin.Context)

	GetUserValue(c *gin.Context)

	TestAuth(c *gin.Context)

	AsciiJSON(c *gin.Context)

	CallBack(c *gin.Context)

	BindJSON(c *gin.Context)

	FormData(c *gin.Context)

	JSON(c *gin.Context)

	PureJSON(c *gin.Context)

	Query(c *gin.Context)

	SecureJSON(c *gin.Context)

	SomeXML(c *gin.Context)

	SomeYAML(c *gin.Context)

	SomeProtoBuf(c *gin.Context)

	Upload(c *gin.Context)

	UploadMulti(c *gin.Context)

	GetFromReader(c *gin.Context)

	TestSecret(c *gin.Context)

	Async(c *gin.Context)

	QueryMap(c *gin.Context)

	GetAndSetCookie(c *gin.Context)

	Schedule(c *gin.Context)

	QueryUser(c *gin.Context)
}
