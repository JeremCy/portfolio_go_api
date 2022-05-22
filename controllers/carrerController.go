package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/models"
)

type CarrerController struct{}

// crud controller for carrer
func (cc *CarrerController) Create(c *gin.Context) {
	var carrer models.Carrer
	c.BindJSON(&carrer)
	models.DB.Create(&carrer)
	c.JSON(http.StatusCreated, gin.H{"message": "Carrer created successfully!", "carrer": carrer})
}

func (cc *CarrerController) FetchAll(c *gin.Context) {
	var carrers []models.Carrer
	models.DB.Find(&carrers)
	c.JSON(http.StatusOK, gin.H{"carrers": carrers})
}

func (cc *CarrerController) FetchOne(c *gin.Context) {
	id := c.Params.ByName("id")
	var carrer models.Carrer
	models.DB.First(&carrer, id)
	c.JSON(http.StatusOK, gin.H{"carrer": carrer})
}

func (cc *CarrerController) Update(c *gin.Context) {
	var carrer models.Carrer
	id := c.Params.ByName("id")
	if models.DB.First(&carrer, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	c.BindJSON(&carrer)
	models.DB.Save(&carrer)
	c.JSON(http.StatusOK, gin.H{"message": "Carrer updated successfully!", "carrer": carrer})
}

func (cc *CarrerController) Delete(c *gin.Context) {
	var carrer models.Carrer
	id := c.Params.ByName("id")
	if models.DB.First(&carrer, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	models.DB.Delete(&carrer)
	c.JSON(http.StatusOK, gin.H{"message": "Carrer deleted successfully!"})
}
