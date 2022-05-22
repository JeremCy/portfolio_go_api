package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func (uc *AuthController) Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	//crypt user password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "user": user})
}

// Login user using jwt token
func (uc *AuthController) Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	var dbUser models.User
	models.DB.Where("email = ?", user.Email).First(&dbUser)
	if dbUser.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found!"})
		return
	}
	//compare password
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password!"})
		return
	}
	//create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  dbUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while signing the token!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": dbUser})
}
