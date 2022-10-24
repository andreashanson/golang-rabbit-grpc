package main

import (
	"fmt"

	"github.com/andreashanson/golang-rabbit-grpc/internal/config"
	"github.com/andreashanson/golang-rabbit-grpc/internal/publisher"
	"github.com/andreashanson/golang-rabbit-grpc/internal/scheduler"
	"github.com/andreashanson/golang-rabbit-grpc/pkg/rabbit"
)

func main() {
	cfg := config.NewConfig()

	conn, err := rabbit.NewRabbitConnection(cfg.RabbitConfig.AmqpServerUrl)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	rabbitRepo := rabbit.NewRabbit(conn)
	pubSvc := publisher.NewService(rabbitRepo)
	schSvc := scheduler.NewService(pubSvc, cfg.RabbitConfig)

	errChan := schSvc.Start()
	defer close(errChan)
	for {
		for errs := range errChan {
			fmt.Println(errs.Error())
		}
	}

}
