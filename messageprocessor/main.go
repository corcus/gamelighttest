package main

import (
	"fmt"
	"gamelighttest/shared"
)

func main() {
	rmqClient := &shared.RmqClient{
		Uri: "rabbitMQ connection url", // would be read from a config file or aws parameter store
	}
	err := rmqClient.Connect()
	if err != nil {
		panic(err)
	}

	redisRepo := shared.NewRedisRepo("redis uri")

	messages, err := rmqClient.Consume("queue name")
	if err != nil {
		panic(err)
	}

	for message := range messages {
		//unmarshal message.Body and create redisDTO
		fmt.Println(string(message.Body))
		dto := shared.RedisDTO{}
		redisRepo.Store(dto)
	}
}
