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

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userID

	err := db.Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}

func CommentList(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}
	var data []interface{}

	err := db.Model(&models.Comment{}).Preload("User").Preload("Photo").Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for i := range Comments {
		photo := make(map[string]interface{})
		user := make(map[string]interface{})

		user["id"] = Comments[i].User.ID
		user["email"] = Comments[i].User.Email
		user["username"] = Comments[i].User.Username

		photo["id"] = Comments[i].Photo.ID
		photo["title"] = Comments[i].Photo.Title
		photo["caption"] = Comments[i].Photo.Caption
		photo["photo_url"] = Comments[i].Photo.PhotoUrl
		photo["user_id"] = Comments[i].Photo.UserId

		data = append(data, gin.H{
			"id":         Comments[i].ID,
			"message":    Comments[i].Message,
			"photo_id":   Comments[i].PhotoId,
			"user_id":    Comments[i].UserId,
			"created_at": Comments[i].CreatedAt,
			"User":       user,
			"Photo":      photo,
		})
	}

	c.JSON(http.StatusOK, data)
}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	photo := []models.Comment{}
	var data []interface{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(commentId)
	Comment.UserId = uint(userId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).First(&Comment).Error
	res := db.Model(&Comment).Preload("Photo").Where("id = ?", commentId).First(&photo).Error

	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": res.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for i := range photo {
		photos := make(map[string]interface{})

		photos["id"] = photo[i].Photo.ID
		photos["title"] = photo[i].Photo.Title
		photos["caption"] = photo[i].Photo.Caption
		photos["photo_url"] = photo[i].Photo.PhotoUrl
		photos["user_id"] = photo[i].Photo.UserId
		photos["updated_at"] = photo[i].Photo.UpdatedAt

		data = append(data, photos)

	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"id":         Comment.ID,
	// 	"message":    Comment.Message,
	// 	"photo_id":   Comment.PhotoId,
	// 	"user_id":    Comment.UserId,
	// 	"updated_at": Comment.UpdatedAt,
	// })
	c.JSON(http.StatusOK, data)
}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	Comment.ID = uint(commentId)
	Comment.UserId = userId

	err := db.Model(&Comment).Where("id = ?", commentId).Delete(models.Comment{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
