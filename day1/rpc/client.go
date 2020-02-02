package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	//Establish a network connection
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err == nil {
		var pd int
		/*
			func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {

		*/
		err = cli.Call("Panda.Getinfo", 1, &pd)
		if err != nil {
			fmt.Println("Call failed")
		}

		fmt.Println("Received value：", pd)
	}

	fmt.Println("Network connection failed")

}
