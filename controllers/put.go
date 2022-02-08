package controllers

import (
	"net/http"

	"github.com/brennschlus/house_points_server/database"
	"github.com/brennschlus/house_points_server/models"
	"github.com/gin-gonic/gin"
)

func UpdateFaculties(c *gin.Context) {
	var json []models.Faculty

	if c.BindJSON(&json) == nil {

		database.GetDB().Save(&json)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})

	}

}
