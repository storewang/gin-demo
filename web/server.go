package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/storewang/gin-demo/web/config"
	"github.com/storewang/gin-demo/web/controller"
)

type Server struct {
	server *gin.Engine
	config config.Config
}

func NewServer() *Server {
	server := &Server{}
	server.initConfig()
	server.init()
	server.initRoute()

	return server
}

func (s *Server) Run() {
	fmt.Println("listen ", s.config.Server.Host)
	s.server.Run(s.config.Server.Host)
	// srv := &http.Server{
	// 	Addr:    s.config.Server.Host,
	// 	Handler: s.server,
	// }

	// go func() {
	// 	// 服务连接
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// // 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("Shutdown Server ...")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// log.Println("Server exiting")
}

func (s *Server) init() {
	s.server = gin.Default()
}
func (s *Server) initConfig() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./web/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}
	viper.Unmarshal(&s.config)
	fmt.Println("-------------------------")
	fmt.Println(s.config)
	fmt.Println("-------------------------")
}
func (s *Server) initRoute() {
	controllers := controller.GetControllers()

	for _, c := range controllers {
		mthd := c.Method
		if mthd == "GET" {
			s.server.GET(c.Path, c.Handler)
		} else if mthd == "POST" {
			s.server.POST(c.Path, c.Handler)
		} else if mthd == "DELETE" {
			s.server.DELETE(c.Path, c.Handler)
		} else if mthd == "PUT" {
			s.server.PUT(c.Path, c.Handler)
		} else if mthd == "" {
		} else {
			fmt.Println("不支持的请求方法:", mthd)
		}
	}
}
