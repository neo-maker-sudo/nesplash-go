package controller

import (
	"gin/api/database"
	"gin/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init (){
	db = database.GetDataBase()
}

func ArchitecturePhotos(c *gin.Context){
	page, _ := strconv.Atoi(c.Query("page"))
	var photos []model.Photo
	db.Preload("User").Where("category_id = ?", 1).Order("Id desc").Offset(int(page)*12).Limit(12).Find(&photos)
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

func ArchitectureVideos(c *gin.Context){
	page,_ := strconv.Atoi(c.Query("page"))
	var videos []model.Video
	db.Where("category_id = ?", 1).Order("Id asc").Offset(int(page)*12).Limit(12).Find(&videos)
	if len(videos) < 12 {
		c.JSON(http.StatusOK, gin.H{
			"message": videos,
			"nextPage": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": videos,
			"nextPage": page + 1,
		})
	}
}

func ArchitectureContributor(c *gin.Context){
	page, _ := strconv.Atoi(c.Query("page"))
	var users []model.User
	db.Joins("LEFT JOIN Photo on user.id = photo.author_id").Where("category_id = ?", 1).Group("username").Order("total_photos desc").Offset(int(page)*12).Limit(12).Find(&users)
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