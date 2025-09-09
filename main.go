package main

import (
	"wecahtdemo/pkg/server"
)

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
