package controller

import (
	"demo2_backend/common"
	"demo2_backend/dto"
	"demo2_backend/model"
	"demo2_backend/response"
	"demo2_backend/util"
	"github.com/gin-gonic/gin"
)

func Commit(ctx *gin.Context)  {
	DB := common.GetDb()

	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		response.Fail(ctx, gin.H{"err":err}, "参数绑定失败")
		return
	}

	//如果前端传来的
	newArticle := model.Article{
		Title: json["title"].(string),
		Tag:   json["tag"].(string),
		Body: json["body"].(string),
	}

	err2 := DB.Create(&newArticle).Error
	if err2 != nil {
		response.Fail(ctx, gin.H{"err":err2}, "文章创建失败")
		return
	}

	response.Success(ctx, nil, "添加成功")
}

func GetArticleByTitle(ctx *gin.Context) {
	DB := common.GetDb()

	json := make(map[string]interface{})
	err1 := ctx.BindJSON(&json)
	if err1 != nil {
		response.Fail(ctx, gin.H{"err":err1}, "请求失败")
		return
	}

	var article model.Article
	err2 := DB.First(&article, "title = ?", json["title"].(string)).Error
	if err2 != nil {
		response.Fail(ctx, gin.H{"err": err2}, "请求失败")
		return
	}

	response.Success(ctx, gin.H{"article":dto.ToArticleDto(article)},"获取成功")
}

func GetArticleByTag(ctx *gin.Context) {
	DB := common.GetDb()

	json := make(map[string]interface{})
	err1 := ctx.BindJSON(&json)
	if err1 != nil {
		response.Fail(ctx, gin.H{"err":err1}, "请求失败")
		return
	}

	var articles []model.Article
	err2 := DB.Where("tag = ?", json["tag"].(string)).Find(&articles).Error
	if err2 != nil {
		response.Fail(ctx, gin.H{"err": err2}, "请求失败")
		return
	}

	response.Success(ctx, gin.H{"article":dto.ToArticlesDto(articles)},"获取成功")
}

func UploadImg(ctx *gin.Context) {
	file, _ := ctx.FormFile("image")
	name := util.RandomString(10)
	filename := util.GetMd5String(name) + ".png"
	if err2 := ctx.SaveUploadedFile(file, "你想存储的位置+/"+filename); err2 != nil {
		response.Fail(ctx, gin.H{"err": err2}, "请求失败")
		return
	}
	response.Success(ctx, gin.H{"url": "服务器地址+存储位置+/" + filename},"上传成功")
}

func FindAll(ctx *gin.Context)  {
	DB := common.GetDb()

	var articles []model.Article
	err := DB.Find(&articles).Error

	if err != nil {
		response.Fail(ctx, gin.H{"err":err}, "获取失败")
		return
	}

	response.Success(ctx, gin.H{"articles":dto.ToArticlesDto(articles)}, "获取成功")
}