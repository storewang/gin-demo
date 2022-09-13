package main

import (
	"github.com/storewang/gin-demo/web"
	_ "github.com/storewang/gin-demo/web/controller"
)

func main() {
	web.NewServer().Run()
}
