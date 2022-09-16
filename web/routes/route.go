package routes

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/shitou/go-demo-gin/infrastructure/utils"
	"github.com/shitou/go-demo-gin/web/handlers"
)

// 使用自动注册路由
type GroupRoute struct {
	Group  string
	Routes []Route
}
type Route struct {
	path       string        //url路径
	httpMethod string        //http方法 get post
	Method     reflect.Value //方法路由
}

func InitRoute(dbutil utils.DbUtil, k8suril *utils.K8sUtil, server *gin.Engine) {
	server.GET("/", handlers.Index)

	server.GET("/path/:id", handlers.PathGet)
	server.POST("/path", handlers.PathPost)
	server.PUT("/path", handlers.PathPut)
	server.DELETE("/path/:id", handlers.PathDelete)
	server.POST("/pathBind", handlers.PathPostJsonBind)
	server.POST("/pathBindf", handlers.PathPostFormBind)
	server.POST("/PathUpload", handlers.PathUpload)

	userHandler := handlers.NewUserHandler(dbutil.Db)
	rg := server.Group("/u")
	rg.GET("/:id", userHandler.FindById)
	rg.POST("/", userHandler.AddUser)

	K8sHandler := handlers.NewK8sHandler(k8suril)
	rg = server.Group("/k8s")
	rg.GET("/ns", K8sHandler.GetAllNamespace)
	rg.GET("/svc", K8sHandler.GetSvcListByNs)
}
