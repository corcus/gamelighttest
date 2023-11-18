package main

import (
	"encoding/json"
	"gamelighttest/shared"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

func main() {

	rmqClient := &shared.RmqClient{
		Uri: "rabbitMQ connection url", // would be read from a config file or aws parameter store
	}
	err := rmqClient.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/message", func(c *gin.Context) {
		var request IncomingRequest
		err := c.BindJSON(&request)
		if err != nil || request.Validate() == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			return
		}

		marshalledrequest, err := json.Marshal(request)
		if err != nil {
			log.Println(err)
		}
		err = rmqClient.Publish(shared.PublishableMessage{
			Key: "testqueuname",
			Publishing: amqp.Publishing{
				ContentType: "application/json",
				Body:        marshalledrequest,
			},
		})
		if err != nil {
			log.Println(err)
		}
	})

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
