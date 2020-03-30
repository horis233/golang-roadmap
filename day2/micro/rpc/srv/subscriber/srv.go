package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	srv "micro/rpc/srv/proto/srv"
)

type Srv struct{}

func (e *Srv) Handle(ctx context.Context, msg *srv.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *srv.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
