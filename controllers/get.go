package controllers

import (
	"net/http"

	"github.com/brennschlus/house_points_server/database"
	"github.com/brennschlus/house_points_server/models"
	"github.com/gin-gonic/gin"
)

func GetFaculties(c *gin.Context) {
	var faculties []models.Faculty
	database.GetDB().Find(&faculties)

	c.JSON(http.StatusOK, gin.H{"data": faculties})
}
