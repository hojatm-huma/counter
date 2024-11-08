package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const ERROR = "ERROR"

func main() {
	db, err := NewDatabase()
	if err != nil {
		fmt.Println(ERROR, "Could not connect to db", err)
	}
	db.Migrate()

	r := gin.Default()

	r.GET("/counter/", func(c *gin.Context) {
		counter := db.FirstOrCreateCounter()
		c.JSON(200, gin.H{"value": counter.Value})
	})

	r.POST("/counter/", func(c *gin.Context) {
		counter := db.FirstOrCreateCounter()
		counter.Increment()
		db.SaveCounter(counter)

		c.JSON(200, gin.H{"value": counter.Value})
	})

	err = r.Run()
	if err != nil {
		fmt.Println(ERROR, "Could not run web server", err)
	}
}
