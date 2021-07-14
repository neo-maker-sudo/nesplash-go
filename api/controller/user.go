package controller

import (
	"gin/api/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

var user model.User

func Member(c *gin.Context){
	session := sessions.Default(c)
	json := &UserData{}
	// 登入
	if c.Request.Method == "PATCH" {

		jsonErr := c.BindJSON(json)
		if jsonErr != nil {
			return 
		}
		if err := db.Where("email = ?", json.Email).First(&user).Error; err != nil {
			c.JSON(400, gin.H{
				"message": "none exist user",
				"error": true,
			})
		} else {
			if user.Password != json.Password {
				c.JSON(400, gin.H{
					"message": "wrong password",
					"error": true,
				})
			} else {
				session.Set("email", json.Email)
				err := session.Save()
				if err != nil {
					return 
				}
				c.JSON(http.StatusOK, gin.H{
					"ok": true,
				})
			}
		}
	} else if c.Request.Method == "POST" {
		jsonErr := c.BindJSON(json)
		if jsonErr != nil {
			return
		}
		var adduser = model.User{Email: json.Email, Username: json.Username, Password: json.Password}
		if err := db.Create(&adduser).Error; err != nil {
			if err := db.Where("email = ?", adduser.Email).First(&adduser).Error; err == nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "email has already been taken, please use another one",
					"error": true,
				})
			} else {
				if err := db.Where("username = ?", adduser.Username).First(&adduser).Error; err == nil {
					c.JSON(400, gin.H{
						"message": "duplicate username",
						"error":   true,
					})
				}
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"method": user,
				"ok": true,
			})
		}
	} else if c.Request.Method == "DELETE" {
		session.Delete("email")
		err := session.Save()
		if err != nil {
			return 
		}
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	} else {
		sess := session.Get("email")
		if sess != nil {
			db.Where("email = ?", sess).First(&user)
			c.JSON(http.StatusOK, gin.H{
				"id": user.Id,
				"email": user.Email,
				"username": user.Username,
				"role_id": user.RoleId,
				"lock_status": user.LockStatus,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": nil,
			})
		}

	}


}