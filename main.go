package main

import (
	"fmt"
	"net/http"
	"pass/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/generate", func(c *gin.Context) {

		length := 12
		if lenStr := c.PostForm("length"); lenStr != "" {
			fmt.Sscanf(lenStr, "%d", &length)
		}

		includeUppercase := c.PostForm("includeUppercase") == "on"
		includeLowercase := c.PostForm("includeLowercase") == "on"
		includeNumbers := c.PostForm("includeNumbers") == "on"
		includeSpecialChars := c.PostForm("includeSpecialChars") == "on"

		if includeUppercase != true {
			includeUppercase = true
		}
		if includeSpecialChars != true {
			includeSpecialChars = true
		}

		password := helper.GeneratePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars)

		c.HTML(http.StatusOK, "generated_password.html", gin.H{
			"password": password,
		})
	})

	router.Run(":8080")
}
