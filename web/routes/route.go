package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shitou/go-demo-gin/infrastructure/utils"
	"github.com/shitou/go-demo-gin/web/handlers"
)

func InitRoute(dbutil utils.DbUtil, server *gin.Engine) {
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
}
