package main

import (
	"net/http"
	"time"

	"github.com/Abhinav-SK07/recipes-api/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/recipes", NewRecipeHandler)
	r.Run() // listen and serve on
}

var recipes []models.Recipe

func init() {
	recipes = []models.Recipe{}
}

func NewRecipeHandler(c *gin.Context) {
	// This is where you would handle the logic for creating a new recipe
	// For example, you could bind the JSON body to a Recipe struct and save it to a database

	var newRecipe models.Recipe
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newRecipe.ID = xid.New().String()    // Generate a unique ID for the recipe
	newRecipe.PublishedAt = time.Now()   // Set the published date to the current time
	recipes = append(recipes, newRecipe) // Add the new recipe to the in-memory slice
	c.JSON(http.StatusOK, newRecipe)     // Return the newly created recipe as a response

}
