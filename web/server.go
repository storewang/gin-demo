package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shitou/go-demo-gin/infrastructure/po"
	"github.com/shitou/go-demo-gin/infrastructure/utils"
	"github.com/shitou/go-demo-gin/web/config"
	"github.com/shitou/go-demo-gin/web/routes"
	"github.com/spf13/viper"
)

type Server struct {
	server  *gin.Engine
	conf    config.Config
	dbutil  utils.DbUtil
	k8suril *utils.K8sUtil
}

func (s *Server) initConfig() {
	viper.SetConfigName("server")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./web/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}

	err = viper.Unmarshal(&s.conf)
	if err != nil {
		log.Fatal("read config failed: ", err)
	}
	log.Println("read config success .")
}

func (s *Server) getAddress() string {
	c := s.conf
	addr := ""
	if c.Server.Host != "" {
		addr += c.Server.Host
	}
	if c.Server.Port != "" {
		addr += ":" + c.Server.Port
	}
	log.Println("host address: ", addr)
	return addr
}

func NewServer() *Server {
	s := &Server{}
	s.init()

	return s
}

func (s *Server) init() *Server {
	// 初始化配置
	s.initConfig()
	// 初始数据库
	s.dbutil = utils.DbUtil{}
	s.dbutil.Init(&s.conf.Database)
	// 初始化k8sUtil
	s.k8suril = utils.NewK8sUtil(&s.conf)

	// 初始化gin
	s.server = gin.Default()
	// 初始化route
	routes.InitRoute(s.dbutil, s.k8suril, s.server)

	return s
}

func (s *Server) Run() {
	// 初始化表结构
	s.dbutil.Db.AutoMigrate(&po.User{})

	// 启动服务器
	s.server.Run(s.getAddress())
}
