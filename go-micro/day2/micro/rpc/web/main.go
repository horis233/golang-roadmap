package main

import (
	"net/http"

	"github.com/micro/go-log"

	"github.com/horis233/golang-roadmap/day2/micro/grpc/web/handler"
	"github.com/micro/go-web"
)

func main() {
	// create new web service
	//创建1个新的web服务
	service := web.NewService(
		//服务名称
		web.Name("go.micro.web.web"),
		//版本号
		web.Version("latest"),
		//设置服务的端口号
		web.Address(":8080"),
	)

	// initialise service
	//初始化服务
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	//注册当前的web前段页面
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	//注册call的请求
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
	//运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
