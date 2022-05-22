package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/models"
)

//CRUD for user

type UserController struct{}

func (uc *UserController) FetchAll(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) FetchOne(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	models.DB.First(&user, id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) Update(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if models.DB.First(&user, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	c.BindJSON(&user)
	models.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!", "user": user})
}

func (uc *UserController) Delete(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if models.DB.First(&user, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
