package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Language struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Paradigm    string `json:"paradigm"`
	Typing      string `json:"typing"`
	YearCreated int    `json:"year_created"`
}

var languages = []Language{
	{ID: 1, Name: "Python", Paradigm: "Multi-paradigm", Typing: "Dynamic", YearCreated: 1991},
	{ID: 2, Name: "JavaScript", Paradigm: "Event-driven", Typing: "Dynamic", YearCreated: 1995},
	{ID: 3, Name: "Go", Paradigm: "Concurrent", Typing: "Static", YearCreated: 2009},
}

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/languages", getLanguages)
	router.GET("/languages/:id", getLanguage)
	router.POST("/languages", createLanguage)
	router.PUT("/languages/:id", updateLanguage)
	router.DELETE("/languages/:id", deleteLanguage)

	router.Run(":8080")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Programming Languages API is working!"})
}

func getLanguages(c *gin.Context) {
	c.JSON(http.StatusOK, languages)
}

func getLanguage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, language := range languages {
		if language.ID == id {
			c.JSON(http.StatusOK, language)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Language not found"})
}

func createLanguage(c *gin.Context) {
	var newLanguage Language

	if err := c.ShouldBindJSON(&newLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	languages = append(languages, newLanguage)
	c.JSON(http.StatusCreated, newLanguage)
}

func updateLanguage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedLanguage Language

	if err := c.ShouldBindJSON(&updatedLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	for index, language := range languages {
		if language.ID == id {
			languages[index] = updatedLanguage
			c.JSON(http.StatusOK, updatedLanguage)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Language not found"})
}

func deleteLanguage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, language := range languages {
		if language.ID == id {
			languages = append(languages[:index], languages[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Language deleted",
				"deleted": language,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Language not found"})
}
