package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	pd "github.com/horis233/golang-roadmap/day2/myproto"
	"context"
)

type server struct {}

// rpc
// function keyword (object) function name (content sent by the client, content returned to the client) error return value

// grpc
// function keyword (object) function name (cotext, parameters sent by the client) (parameters sent to the client, error)

// Sayhello is a greeting function
func (*server)Sayhello(ctx context.Context, in *pd.HelloReq) (out * pd.HelloRsp, err error){

	return  &pd.HelloRsp{Msg:"hello"+in.Name},nil
}
// Sayname is a  service that says name
func (*server)Sayname(ctx context.Context, in *pd.NameReq) ( out *pd.NameRsp, err error){

	return &pd.NameRsp{Msg:in.Name+"Morning"},nil
}

func main() {
	//Establish network
	ln ,err :=net.Listen("tcp",":10086")
	if err !=nil{
		fmt.Println("Network error",err)
	}

	//Establish grpc service
	srv:=grpc.NewServer()

	//Register server
	pd.RegisterHelloserverServer(srv,&server{})

	//Wait for server ready
	err=srv.Serve(ln)
	if err!=nil {
		fmt.Println("Network error",err)
	}

}
