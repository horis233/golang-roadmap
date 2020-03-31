package main

import (
        "github.com/micro/go-log"
	"net/http"

        "github.com/micro/go-web"
        "day2/micro/rpc/web/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.web"),
                web.Version("latest"),
                web.Address(":8080"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}