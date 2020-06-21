package main

import (
	"google.golang.org/grpc"
	"fmt"
	pd "github.com/horis233/golang-roadmap/day2/myproto"
	"context"
)

func main() {
	//Client connects to the server
	conn ,err :=grpc.Dial("127.0.0.1:10086",grpc.WithInsecure())
	if err!=nil {
		fmt.Println("Network err",err)
	}
	//Close network connection
	defer  conn.Close()


	//Get grpc client handles
	c:=pd.NewHelloserverClient(conn)

	//Calling functions through client handles
	re ,err :=c.Sayhello(context.Background(),&pd.HelloReq{Name:"Panda"})
	if err!=nil {
		fmt.Println("sayhello service call failed")
	}
	fmt.Println("service call failed",re.Msg)



	re1 ,err :=c.Sayname(context.Background(),&pd.NameReq{Name:"Jiaming"})
	if err !=nil{
		fmt.Println("say name service call failed")
	}
	fmt.Println("service call failed",re1.Msg)

}