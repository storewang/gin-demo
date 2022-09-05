package web

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

//路由结构体
type Route struct {
	path       string        //url路径
	httpMethod string        //http方法 get post
	Method     reflect.Value //方法路由
}

//路由集合
var Routes = []Route{}

func NewRoute(path, httpMethod string, method reflect.Value) *Route {
	return &Route{path: path, httpMethod: httpMethod, Method: method}
}
func InitRouter(r *gin.Engine) {
	//绑定基本路由，访问路径：/User/List
	Bind(r)
}

func getMethodAndAction(action string) (string, string) {
	action = strings.ToLower(action)
	index := strings.Index(action, "get")
	httpMethod := ""
	if index > -1 {
		httpMethod = "GET"
		action = strings.Replace(action, "get", "", 1)
	}
	index = strings.Index(action, "post")
	if index > -1 {
		httpMethod = "POST"
		action = strings.Replace(action, "post", "", 1)
	}
	index = strings.Index(action, "put")
	if index > -1 {
		httpMethod = "PUT"
		action = strings.Replace(action, "put", "", 1)
	}
	index = strings.Index(action, "delete")
	if index > -1 {
		httpMethod = "DELETE"
		action = strings.Replace(action, "delete", "", 1)
	}

	return httpMethod, action
}

// 注册控制器
func RegisterRoute(rt *Route) bool {
	Routes = append(Routes, *rt)
	return true
}

//注册控制器
func Register(module string, controller interface{}) bool {
	v := reflect.ValueOf(controller)
	//遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name

		// 获取path,method
		action = strings.ToLower(action)
		httpMethod, action := getMethodAndAction(action)
		path := "/" + module + "/" + action

		if httpMethod != "" {
			route := Route{path: path, Method: method, httpMethod: httpMethod}
			Routes = append(Routes, route)
		}
	}
	fmt.Println("Routes=", Routes)
	return true
}

//绑定路由 m是方法GET POST等
//绑定基本路由
func Bind(e *gin.Engine) {
	for _, route := range Routes {

		if route.httpMethod == "GET" {
			e.GET(route.path, match(route))
		}

		if route.httpMethod == "POST" {
			e.POST(route.path, match(route))
		}

		if route.httpMethod == "PUT" {
			e.PUT(route.path, match(route))
		}

		if route.httpMethod == "DELETE" {
			e.DELETE(route.path, match(route))
		}
	}
}

//根据path匹配对应的方法
func match(route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(route.httpMethod) < 3 {
			return
		}

		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context

			route.Method.Call(arguments)
		}
	}
}
