// بسم الله الرحمن الرحيم
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", HelloHandler)
	v1.GET("/books/:id/:title", booksHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", postBookHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Muhdan Fyan Syah",
		"bio":  "Koran Teacher",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Halo dunia",
		"subtitle": "Halo kehidupan",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{"id": title, "price": price})
}

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	// Price int    `json:"price" binding:"required,number"`
}

func postBookHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {

		// for _, e := range err.(validator.ValidationErrors) {
		// 	errorMessage := fmt.Sprintf("Error on field %s, condtion: %s", e.Field(), e.ActualTag())
		// 	c.JSON(http.StatusBadRequest, errorMessage)
		// 	return
		// }

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condtion: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
