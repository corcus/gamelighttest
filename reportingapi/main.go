package main

import (
	"encoding/json"
	"gamelighttest/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	redisRepo := shared.NewRedisRepo("redis uri")

	router := gin.Default()

	router.GET("/message/list", func(c *gin.Context) {
		sender := c.Query("sender")
		receiver := c.Query("receiver")

		redisDTOs, err := redisRepo.Get(sender, receiver)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}

		marshalledDTOs, err := json.Marshal(redisDTOs)
		c.JSON(http.StatusOK, gin.H{"list": marshalledDTOs})
	})

	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}
}
