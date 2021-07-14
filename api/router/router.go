package router

import (
	"gin/api/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	store := cookie.NewStore([]byte("secretkey"))
	router.Use(sessions.Sessions("sessionId", store))
	v1 := router.Group("/api")
	{
		v1.GET("/photos", controller.HomePhotos)
		v1.GET("/photos/search", controller.SearchPhoto)
		v1.GET("/users/search", controller.SearchUser)
		v1.GET("/collected_photo_id", controller.CollectedPhotoId)
		v1.GET("/user", controller.Member)
		v1.POST("/user", controller.Member)
		v1.PATCH("/user", controller.Member)
		v1.DELETE("/user", controller.Member)
	}

	v2 := router.Group("/architecture/api")
	{
		v2.GET("/photos", controller.ArchitecturePhotos)
		v2.GET("/videos", controller.ArchitectureVideos)
		v2.GET("/contributor", controller.ArchitectureContributor)
	}

	v3 := router.Group("/athletics/api")
	{
		v3.GET("/photos", controller.AthleticsPhotos)
		v3.GET("/videos", controller.AthleticsVideos)
		v3.GET("/contributor", controller.AthleticsContributor)
	}
	return router
}

