package main

import (
	"fmt"

	"github.com/raven0520/btc/app"
	"github.com/raven0520/btc/router"
)

var env = "./config/"

func main() {
	err := app.InitModule(env, []string{"base", "node"})
	if err != nil {
		fmt.Printf("Failed Init %s", err) // 输出错误信息到控制台
		return
	}
	if err = app.InitConsulServer(); err != nil {
		fmt.Println(err)
		return
	}
	server := router.NewGrpcServer()
	server.GrpcRegister()
	server.GrpcServerStart()      // Start Grpc Service
	defer server.GrpcUnregister() // Deregister Grpc Service
}
