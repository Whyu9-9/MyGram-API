package controllers

import (
	"mygram-api/database"
	"mygram-api/helpers"
	"mygram-api/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	if err := db.Debug().Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_email_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "Email already exists",
			})

			return
		}

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_username_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "Username already exists",
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	originalPassword := user.Password
	if err := db.Debug().Where("email = ?", user.Email).First(&user).Take(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "User not found",
		})

		return
	}

	if isValid := helpers.CheckPasswordHash([]byte(user.Password), []byte(originalPassword)); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})

		return
	}

	jwt := helpers.GenerateToken(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})

}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	user := models.User{}

	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid user id",
		})

		return
	}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	user.ID = userID

	err = db.Debug().Model(&user).Where("id = ?", userid).Updates(&user).First(&user).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_email_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "Email already exists",
			})

			return
		}

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_username_key\"") {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "Username already exists",
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"updated_at": user.UpdatedAt,
		"age":        user.Age,
	})

}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	user := models.User{}

	userID := uint(userData["id"].(float64))

	user.ID = userID

	err := db.Model(&user).Where("id = ?", userID).Delete(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})

}
