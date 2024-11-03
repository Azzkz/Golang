package handlers

import (
	"GO3/models"
	"GO3/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var users = []models.User{}
var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Регистрация нового пользователя
func Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем уникальность имени пользователя
	for _, user := range users {
		if user.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
	}

	// Сохраняем нового пользователя
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	utils.Logger.Println("New user registered:", newUser.Username)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Вход пользователя и генерация JWT токена
func Login(c *gin.Context) {
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundUser *models.User
	for _, user := range users {
		if user.Username == loginUser.Username && user.Password == loginUser.Password {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: foundUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	utils.Logger.Println("User logged in:", foundUser.Username)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
