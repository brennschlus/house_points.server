package main

import (
	"net/http"

	"github.com/brennschlus/house_points_server/database"
	"github.com/brennschlus/house_points_server/models"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "dangerous zone, idk how servers works!!!",
		})
	})
	r.GET("/faculties", getFaculties)
	r.Run()
}

func getFaculties(c *gin.Context) {
	var faculties []models.Faculty
	database.GetDB().Find(&faculties)

	c.JSON(http.StatusOK, gin.H{"data": faculties})
}
