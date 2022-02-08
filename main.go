package main

import (
	"github.com/brennschlus/house_points_server/controllers"
	"github.com/brennschlus/house_points_server/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API для программы Факультетские часы Хогвартса",
		})
	})
	r.GET("/faculties", controllers.GetFaculties)
	r.PUT("/faculties", controllers.UpdateFaculties)

	r.Run()
}
