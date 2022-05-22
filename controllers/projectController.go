package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/models"
)

// CRUD operations for project
type ProjectController struct{}

// Create a new project
func (pc *ProjectController) Create(c *gin.Context) {
	var project models.Project
	c.BindJSON(&project)
	models.DB.Create(&project)
	c.JSON(http.StatusCreated, gin.H{"message": "Project created successfully!", "project": project})
}

// Fetch all projects
func (pc *ProjectController) FetchAll(c *gin.Context) {
	var projects []models.Project
	models.DB.Find(&projects)
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// Fetch a project with an ID
func (pc *ProjectController) FetchOne(c *gin.Context) {
	id := c.Params.ByName("id")
	var project models.Project
	models.DB.First(&project, id)
	c.JSON(http.StatusOK, gin.H{"project": project})
}

// Update a project with an ID
func (pc *ProjectController) Update(c *gin.Context) {
	var project models.Project
	id := c.Params.ByName("id")
	if models.DB.First(&project, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	c.BindJSON(&project)
	models.DB.Save(&project)
	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully!", "project": project})
}

// Delete a project with an ID
func (pc *ProjectController) Delete(c *gin.Context) {
	var project models.Project
	id := c.Params.ByName("id")
	if models.DB.First(&project, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	models.DB.Delete(&project)
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully!"})
}
