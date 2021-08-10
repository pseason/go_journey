package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由
var Handlers []Route

// http method
type HttpMethod int

const (
	GET HttpMethod = iota + 1
	POST
	DELETE
	PUT
	PATCH
)

// 自定义route结构
type Route struct {
	Url     string
	Method  HttpMethod
	Handler func(c *gin.Context)
}

func createRoute(url string, method HttpMethod, handler func(c *gin.Context)) Route {
	return Route{Url: url, Method: method, Handler: handler}
}

// 加载全部路由
func init() {
	Handlers = make([]Route, 0)
	// hello
	Handlers = append(Handlers, createRoute("/hello", GET, Hello))
	// studnet
	Handlers = append(Handlers, createRoute("/student/find/:id", GET, FindStudentById))
	Handlers = append(Handlers, createRoute("/student/findAll", GET, FindAllStudent))
	Handlers = append(Handlers, createRoute("/student/insert", POST, InsertStudent))
	Handlers = append(Handlers, createRoute("/student/update", POST, UpdateStudent))
	Handlers = append(Handlers, createRoute("/student/delete/:id", GET, DeleteStudent))
}

// gin cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
