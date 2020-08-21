package main

import (
	"github.com/St0iK/go-web-docker/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"alanis":  "jimstoik13",
		})
	})

	r.GET("/animal/:name", func(c *gin.Context) {
		animal, err := database.GetAnimal(c.Param("name"))
		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, animal)
	})
	r.Run(":3000")
}
