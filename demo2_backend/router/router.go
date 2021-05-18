package router

import (
	"demo2_backend/controller"
	"demo2_backend/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.POST("/api/article/uploadArticle", controller.UploadArticle) // ok
	r.POST("/api/article/uploadImg", controller.UploadImg) // ok
	r.POST("/api/article/getAllTags", controller.GetAllTags) // ok
	r.POST("/api/article/getAllArticles", controller.GetAllArticles) // ok
	r.POST("/api/article/getRelationT", controller.GetRelationT) // ok
	r.POST("/api/article/getRelationA", controller.GetRelationA) // ok
	r.POST("/api/article/getArticlesByRelation", controller.GetArticlesByRelation) // ok
	r.POST("/api/article/getArticleByTitle", controller.GetArticleByTitle) // ok
	r.POST("/api/article/delArticle", controller.DelArticle) // ok
	r.POST("/api/article/updateArticle", controller.UpdateArticle) // ok
	r.POST("/api/article/createTagName", controller.CreateTagName) // ok
	r.POST("/api/article/delTagName", controller.DelTagName) // ok
	return r
}