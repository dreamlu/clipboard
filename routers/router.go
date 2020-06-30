// @author  dreamlu
package routers

import (
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/util/str"
	"github.com/gin-gonic/gin"
	"github/dreamlu/clipboard/controllers"
	"net/http"
	"strings"
)

var Router = SetRouter()

var V = Router.Group("/api/v1")

func SetRouter() *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	str.MaxUploadMemory = router.MaxMultipartMemory
	//router.Use(CorsMiddleware())

	// 过滤器
	//router.Use(Filter())

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//组的路由,version
	v1 := router.Group("/api")
	{
		v := v1

		v.GET("/clip", controllers.Clip)
	}
	//不存在路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"msg":    "接口不存在->('.')/请求方法不存在",
		})
	})
	return router
}

// 登录失效验证
func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == "GET" {
			c.Next()
			return
		}
		path := c.Request.URL.String()


		if !strings.Contains(path, "login") && !strings.Contains(path, "/static/file") {
			_, err := c.Cookie("uid") // may be use session
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, result.MapNoAuth)
				return
			}
		}
	}
}

// 处理跨域请求,支持options访问
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		//fmt.Println(method)
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		c.Header("Access-Control-Allow-Credentials", "true")
//
//		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		// 处理请求
//		c.Next()
//	}
//}
