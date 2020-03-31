package main

import (
	"github.com/horis233/golang-roadmap/day2/micro/grpc/srv/handler"
	example "github.com/horis233/golang-roadmap/day2/micro/grpc/srv/proto/example"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	//01加载micro 关于grpc的插件
	"github.com/micro/go-grpc"
)

func main() {
	// New Service
	//创建服务
	/*	service := micro.NewService(
		//服务名
		micro.Name("go.micro.srv.srv"),
		//版本号
		micro.Version("latest"),
	)*/
	//02 将之前的micro缓存grpc
	service := grpc.NewService(
		//服务名
		micro.Name("go.micro.srv.srv"),
		//版本号
		micro.Version("latest"),
	)

	// Initialise service
	//初始化我们的服务
	service.Init()

	// Register Handler
	//注册我们的服务
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//注册结构体 源自于Subscriber
	//micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//注册一个方法 源自于Subscriber
	//micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
