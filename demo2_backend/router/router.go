package router

import (
	"demo2_backend/controller"
	"demo2_backend/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.POST("/api/article/commit", controller.Commit)
	r.POST("/api/article/getArticleByTitle", controller.GetArticleByTitle)
	r.POST("/api/article/getArticleByTag", controller.GetArticleByTag)
	//r.POST("/api/article/uploadImg", controller.UploadImg)
	r.GET("/api/article/findAll", controller.FindAll)
	return r
}


