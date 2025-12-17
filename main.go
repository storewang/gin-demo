package main

import (
	"wechatdemo/pkg/server"
)

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
