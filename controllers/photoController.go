package controllers

import (
	"mygram-api/database"
	"mygram-api/helpers"
	"mygram-api/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func PhotoCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserId,
		"created_at": Photo.CreatedAt,
	})
}

func PhotoGetAll(c *gin.Context) {
	db := database.GetDB()
	var Photos []models.Photo

	var data []interface{}

	err := db.Debug().Preload("User").Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for i := range Photos {
		photo := make(map[string]interface{})
		user := make(map[string]interface{})

		user["email"] = Photos[i].User.Email
		user["username"] = Photos[i].User.Username

		photo["id"] = Photos[i].ID
		photo["title"] = Photos[i].Title
		photo["caption"] = Photos[i].Caption
		photo["photo_url"] = Photos[i].PhotoUrl
		photo["user_id"] = Photos[i].UserId
		photo["created_at"] = Photos[i].CreatedAt
		photo["updated_at"] = Photos[i].UpdatedAt
		photo["User"] = user

		data = append(data, photo)
	}

	c.JSON(http.StatusOK, data)

}

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userId
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserId,
		"updated_at": Photo.UpdatedAt,
	})
}

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	Photo.UserId = userId
	Photo.ID = uint(photoId)

	err := db.Debug().Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
