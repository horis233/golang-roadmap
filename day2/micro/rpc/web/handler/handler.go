package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	//倒入的是案例的proto
	//example "github.com/micro/examples/template/srv/proto/example"
	"fmt"

	example "github.com/horis233/golang-roadmap/day2/micro/grpc/srv/proto/example"
	"github.com/micro/go-grpc"
)
//（传出，传入）
func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	//创建一个map
	var request map[string]interface{}
	//r.body 传入的数据 解码到  request map里面
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("前段发送给过来的json转化为map打印",request["name"])

	//创建1个grpc的服务返回1个句柄
	server :=grpc.NewService()
	//初始化
	server.Init()


	// call the backend service
	//调用服务 返回句柄
	//exampleClient := example.NewExampleService("go.micro.srv.srv", client.DefaultClient)
	exampleClient := example.NewExampleService("go.micro.srv.srv", server.Client())

	//通过句柄 调用call函数
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("服务端做完处理 会传的数据信息",rsp.Msg)

	//接收 服务调用的返回信息 创建成为map
	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}


	// encode and write the response as json
	//将 response 转化为json 发送给前段
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}