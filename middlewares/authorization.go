package middlewares

import (
	"mygram-api/database"
	"mygram-api/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProfileAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		user := models.User{}

		if err := db.Debug().Where("id = ?", userID).First(&user).Take(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			})

			return
		}

		if user.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this resource",
			})

			return
		}
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			})

			return
		}

		if Photo.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this resource",
			})

			return
		}
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			})

			return
		}

		if Comment.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this resource",
			})

			return
		}
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			})
			return
		}

		if SocialMedia.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this resource",
			})

			return
		}
	}
}
