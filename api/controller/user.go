package controller

import (
	"fmt"
	"gin/api/model"
	"gin/api/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email string `json:"email"`
	Bio string `json:"bio"`
	Location string `json:"location"`
}

var user model.User
var photos []model.Photo

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
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "none exist user",
				"error": true,
			})
		} else {
			if user.Password != json.Password {
				c.JSON(http.StatusBadRequest, gin.H{
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
		var adduser = model.User{
			Email: json.Email,
			Username: json.Username,
			Password: json.Password,
			ProfileImage: "https://dkn8b9qqzonkk.cloudfront.net/profile_pics/default.jpg",
			RoleId: 2,
			MethodId: 1,
		}
		if err := db.Create(&adduser).Error; err != nil {
			if err := db.Where("email = ?", adduser.Email).First(&adduser).Error; err == nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "email has already been taken, please use another one",
					"error": true,
				})
			} else {
				if err := db.Where("username = ?", adduser.Username).First(&adduser).Error; err == nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "duplicate username",
						"error":   true,
					})
				}
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
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

func PublicPageApi(c *gin.Context){
	// 外建設定導致這個api的photos產出user資料，但此api內容主要是以另一條query去做查詢，還想不到辦法如何將photos裡面的users弄掉
	page, _ := strconv.Atoi(c.Query("page"))
	userId := c.Param("userId")
	db.Where("id = ?", userId).First(&user)
	db.Where("author_id = ?", userId).Order("Id asc").Offset(int(page)*12).Limit(12).Find(&photos)
	if len(photos) < 12 {
		c.JSON(http.StatusOK, gin.H{
			"message": photos,
			"user": user,
			"nextPage": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": photos,
			"user": user,
			"nextPage": page + 1,
		})
	}
}

func PersonDataApi(c *gin.Context) {
	session := sessions.Default(c)
	sess := session.Get("email")
	if sess != nil {
		db.Preload("Method").Where("email = ?", sess).First(&user)
		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": []int{},
		})
	}
}

func ChangePassword(c *gin.Context) {
	session := sessions.Default(c)
	sess := session.Get("email")
	json := &UserData{}
	jsonErr := c.BindJSON(json)
	if jsonErr != nil {
		return
	}
	if sess != nil {
		if json.ConfirmPassword != json.Password {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"message": "password and confirm_password not same",
			})
			return
		}

		if err := db.Model(&user).Where("email = ?", sess).Update("password", json.Password).Error; err != nil {
			log.Fatalln("Error happening on ChangePassword function : ", err)
		}

		if user.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "none exist user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	} else {
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func ChangeBio(c *gin.Context) {
	session := sessions.Default(c)
	sess := session.Get("email")
	json := &UserData{}
	jsonErr := c.BindJSON(json)

	if jsonErr != nil {
		return
	}
	if sess != nil {
		if err := db.Model(&user).Where("email = ?", sess).Update("bio", json.Bio).Error; err != nil {
			log.Fatalln("Error happening on ChangeBio function : ", err)
		}

		if user.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "none exist user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	} else {
		// redirect index page
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func ChangeUsername(c *gin.Context) {
	session := sessions.Default(c)
	sess := session.Get("email")
	json := &UserData{}
	jsonErr := c.BindJSON(json)
	if jsonErr != nil {
		return
	}
	if sess != nil {
		if err := db.Model(&user).Where("email = ?", sess).Update("username", json.Username).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "username already be taken, change another one",
			})
			return
		}

		if user.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "none exist user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	} else {
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func ChangeLocation(c *gin.Context) {
	session := sessions.Default(c)
	sess := session.Get("email")
	json := &UserData{}
	jsonErr := c.BindJSON(json)
	if jsonErr != nil {
		return
	}
	if sess != nil {
		if err := db.Model(&user).Where("email = ?", sess).Update("location", json.Location).Error; err != nil {
			log.Fatalln("Error happening on ChangeLocation function : ", err)
		}

		if user.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "none exist user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	} else {
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func DeleteAccount(c *gin.Context) {
	session := sessions.Default(c)
	sess := session.Get("email")
	if sess != nil {
		if err := db.Where("email = ?", sess).Delete(&user).Error; err != nil {
			log.Fatalln("Error happening on DeleteAccount function : ", err)
		}

		if user.Id == 0 {
			c.JSON(400, gin.H{
				"error": "none exist user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	} else {
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func UploadProfileImage(c *gin.Context) {
	session := sessions.Default(c)
	sess := session.Get("email")
	if sess != nil {
		db.Where("email = ?", sess).First(&user)
		if user.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : "none exist user",
			})
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		extension := filepath.Ext(file.Filename)
		randomHex, err := util.RandomHex(8)
		fmt.Println(randomHex)
		if err != nil {
			fmt.Println("Hexing error", err)
		}
		cleanFn := randomHex + extension

		if err := c.SaveUploadedFile(file, "static/profile_pics/" + cleanFn); err != nil {
			fmt.Println("Saving image into local directory error : ", err)
		}
		// 尚未連接將檔案上傳至s3，但先將CDN路徑跟檔案名稱合併
		user.ProfileImage = "https://dkn8b9qqzonkk.cloudfront.net/profile_pics/" + cleanFn

		db.Save(&user)
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"message": cleanFn,
		})
	} else {
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func UploadPublicImage(c *gin.Context) {

}

func DeletePublicImage(c *gin.Context) {

}

func PersonalPhotoId(c *gin.Context) {

}