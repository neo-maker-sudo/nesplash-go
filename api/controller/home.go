package controller

import (
	"github.com/gin-contrib/sessions"
	"log"

	//"encoding/json"
	//"fmt"
	"gin/api/database"
	"gin/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB

func init(){
	db = database.GetDataBase()
}


func HomePhotos(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	var photos []model.Photo
	db.Preload("User").Order("Id desc").Offset(page*12).Limit(12).Find(&photos)
	if len(photos) < 12 {
		c.JSON(http.StatusOK, gin.H{
			"message": photos,
			"nextPage": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": photos,
			"nextPage": page + 1,
		})
	}
}

func SearchPhoto(c *gin.Context) {
	q := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page"))
	var photos []model.Photo
	err := db.Preload("User").Where("description LIKE ?", "%"+q+"%").Order("Id desc").Offset(page*12).Limit(12).Find(&photos)
	if err.Error != nil {
		log.Fatalln("Error happening on SearchPhoto function : ", err)
	}

	if len(photos) < 12 {
		c.JSON(http.StatusOK, gin.H{
			"message": photos,
			"nextPage": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": photos,
			"nextPage": page + 1,
		})
	}
}

func SearchUser(c *gin.Context) {
	q := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page"))
	var users []model.User
	if err := db.Where("username LIKE ?", "%"+q+"%").Order("Id desc").Offset(page*12).Limit(12).Find(&users).Error; err != nil {
		log.Fatalln("Error happening on SearchUser function : ", err)
	}
	if len(users) < 12 {
		c.JSON(http.StatusOK, gin.H{
			"message": users,
			"nextPage": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": users,
			"nextPage": page + 1,
		})
	}

}

func CollectedPhotoId(c *gin.Context){
	session := sessions.Default(c)
	sess := session.Get("email")
	if sess != nil {
		var res []int
		var user model.User
		var collections []model.Collection
		if err := db.Where("email = ?", sess).First(&user).Error; err != nil {
			log.Fatalln("(search user) Error happening on CollectionPhotoId function : ", err)
		}
		db.Where("collector_id = ?", user.Id).Order("collected_id asc").Find(&collections)
		for _, val := range collections {
			res = append(res, val.CollectedId)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": []int{},
		})
	}
}