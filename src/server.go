package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func server() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tree", func (c *gin.Context)  {
		cmd := exec.Command("tree")
		output, err := cmd.Output()
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println(string(output))
	})
	r.Run(":5000") // listen and serve on 0.0.0.0:8080
}