package subscriber

import (
	"context"

	"github.com/micro/go-log"

	example "github.com/horis233/golang-roadmap/day2/micro/rpc/srv/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	//处理收到的消息
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	//收到的消息函数
	log.Log("Function Received message: ", msg.Say)
	return nil
}
