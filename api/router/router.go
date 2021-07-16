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
	category := controller.New()
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
		v1.GET("/public/:userId", controller.PublicPageApi)
		v1.GET("/account/data", controller.PersonDataApi)
	}

	apiUser := router.Group("/api/user")
	{
		apiUser.POST("/change-bio", controller.ChangeBio)
		apiUser.POST("/change-password", controller.ChangePassword)
		apiUser.POST("/change-username", controller.ChangeUsername)
		apiUser.POST("/change-location", controller.ChangeLocation)
		apiUser.DELETE("/delete-account", controller.DeleteAccount)
		apiUser.POST("/upload-profile-image", controller.UploadProfileImage)
		apiUser.POST("/upload-public-image", controller.UploadPublicImage)
		apiUser.POST("/delete-public-image", controller.DeletePublicImage)
		apiUser.GET("/personal-photos", controller.PersonalPhotoId)
	}

	v2 := router.Group("/architecture/api")
	{
		v2.GET("/photos", controller.CategoryPhotos(category.Architecture))
		v2.GET("/videos", controller.CategoryVideos(category.Architecture))
		v2.GET("/contributor", controller.CategoryContributor(category.Architecture))
	}
	v3 := router.Group("/athletics/api")
	{
		v3.GET("/photos", controller.CategoryPhotos(category.Athletics))
		v3.GET("/videos", controller.CategoryVideos(category.Athletics))
		v3.GET("/contributor", controller.CategoryContributor(category.Athletics))
	}
	v4 := router.Group("/foodie/api")
	{
		v4.GET("/photos", controller.CategoryPhotos(category.Foodie))
		v4.GET("/videos", controller.CategoryVideos(category.Foodie))
		v4.GET("/contributor", controller.CategoryContributor(category.Foodie))
	}
	v5 := router.Group("/nature/api")
	{
		v5.GET("/photos", controller.CategoryPhotos(category.Nature))
		v5.GET("/videos", controller.CategoryVideos(category.Nature))
		v5.GET("/contributor", controller.CategoryContributor(category.Nature))
	}
	v6 := router.Group("/people/api")
	{
		v6.GET("/photos", controller.CategoryPhotos(category.People))
		v6.GET("/videos", controller.CategoryVideos(category.People))
		v6.GET("/contributor", controller.CategoryContributor(category.People))
	}
	v7 := router.Group("/travel/api")
	{
		v7.GET("/photos", controller.CategoryPhotos(category.Travel))
		v7.GET("/videos", controller.CategoryVideos(category.Travel))
		v7.GET("/contributor", controller.CategoryContributor(category.Travel))
	}
	return router
}

