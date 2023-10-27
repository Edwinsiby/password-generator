package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func generatePassword(length int, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars bool) string {
	charset := ""
	if includeUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeLowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if includeNumbers {
		charset += "0123456789"
	}
	if includeSpecialChars {
		charset += "!@#$%^&*()-_=+[]{}|;:'\",.<>?/\\"
	}

	password := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

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

		password := generatePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars)

		c.HTML(http.StatusOK, "generated_password.html", gin.H{
			"password": password,
		})
	})

	router.Run(":8080")
}
