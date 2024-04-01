package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/dailytask", DailyTask)
	router.Run(":2345")
}
