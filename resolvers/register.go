package resolvers

import (
	"fmt"
	"log"
	"net/http"

	logger "github.com/sahlannasution/xnews-xapiens-backend/log"
	"github.com/sahlannasution/xnews-xapiens-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Users struct
type User struct {
	ID       string `gorm:"primarykey" json:"id"`
	Email    string `gorm:"primarykey" json:"email"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}

// Register func
func (ResolverDB *ResolverDB) Register(c *gin.Context) {
	var (
		users  models.Users
		result gin.H
	)
	if err := c.Bind(&users); err != nil || users.Email == "" || users.Password == "" || users.Fullname == "" {
		e := "Field Email, Password, Fullname is required!"
		result = gin.H{
			"status":  "bad request",
			"message": e,
		}
		fmt.Println("Field Email, Password, Fullname is required!")
		c.JSON(http.StatusBadRequest, result)

		logger.Sentry(err) // push log error ke sentry
	} else {
		users.Roles = 4
		encrypt, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			logger.Sentry(err)
		}
		users.Password = string(encrypt)
		if res := ResolverDB.DB.Create(&users); res.Error != nil {
			err := res.Error
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "bad request",
				"message": err.Error(),
			})

			logger.Sentry(err) // push log error ke sentry
		} else {
			result = gin.H{
				"status":  "success",
				"message": "Sucessfully Register!",
				"data": map[string]interface{}{
					"id":       users.ID,
					"email":    users.Email,
					"fullName": users.Fullname,
					"role":     users.Roles,
				},
			}
			c.JSON(http.StatusOK, result)
		}
	}
}
