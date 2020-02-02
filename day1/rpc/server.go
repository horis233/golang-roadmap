package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

/*
- the method's type is exported.
- the method is exported.
- the method has two arguments, both exported (or builtin) types.
- the method's second argument is a pointer.
- the method has return type error.
func (t *T) MethodName(argType T1, replyType *T2) error
*/

type Panda int

// Function keyword (object) Function name (content sent from the peer, returned to the peer) Error return value
func (*Panda) Getinfo(argType int, replyType *int) error {

	fmt.Println("The content sent by the peer is:", argType)

	//Modify content value
	*replyType = argType + 1

	return nil
}

func main() {

	//Instantiate a class as an object
	pd := new(Panda)
	//The server registers an object
	err := rpc.Register(pd)
	if err != nil {
		fmt.Println("Register Error")
	}
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("Network Error")
	}
	err = http.Serve(ln, nil)
	if err != nil {
		fmt.Println("Server Error")
	}
}
