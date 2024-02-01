package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getNextNumbers(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		panic(err)
	}

	nextNumbers := make([]int, 0)
	nextNumbers = append(nextNumbers, number)

	for i := 0; i < 100; i++ {
		if number%2 == 0 {
			number = number / 2
		} else {
			number = 3*number + 1
		}
		nextNumbers = append(nextNumbers, number)
		if number == 1 {
			break
		}
	}
	c.IndentedJSON(http.StatusCreated, nextNumbers)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	router.GET("/next/:number", getNextNumbers)

	router.Run(":" + port)
}
