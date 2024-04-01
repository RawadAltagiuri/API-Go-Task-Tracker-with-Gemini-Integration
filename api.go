package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type DailyTaskgist struct {
	Text string `json:"text"`
}

// this function will take in a parameter from a post request, and it will pass the string parameter to the GenerateRespone function in gemini.go
func DailyTask(c *gin.Context) {
	var dailyTaskgist DailyTaskgist
	c.BindJSON(&dailyTaskgist)
	a := GenerateRespone(dailyTaskgist.Text)
	fmt.Println(a)
	//remove "-" from the strings in the array, this depends on your prompt so you might not need this
	for i := 0; i < len(a); i++ {
		a[i] = a[i][2:]
	}
	// append the strings to the google sheet, the sheet id is the first parameter, and the array of strings is the second parameter, the sheet id is set as an environment variable
	appendToSheet("YOUR GOOGLE SHEET ID", a)
	c.IndentedJSON(200, a)
}
