package controller

import (
	"demo2_backend/common"
	"demo2_backend/dto"
	"demo2_backend/model"
	"demo2_backend/response"
	"demo2_backend/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//上传
func UploadArticle(ctx *gin.Context)  {
	DB := common.GetDb()

	fjson := make(map[string]interface{})
	err1 := ctx.BindJSON(&fjson)
	if err1 != nil {
		response.Fail(ctx, gin.H{"err":err1}, "请求失败")
		return
	}
	title := fjson["title"].(string)
	body := fjson["body"].(string)
	describe := fjson["describe"].(string)
	writeTime := fjson["writeTime"].(string)
	readingTime := fjson["readingTime"].(string)
	backgroundImage := fjson["backgroundImage"].(string)
	tagArr := fjson["tagName"].([]interface{})
	//数据校验
	if len(title) == 0 || len(body) == 0 || len(describe) == 0 || len(writeTime) == 0 || len(readingTime) == 0 || len(backgroundImage) == 0{
		response.Fail(ctx, nil, "有空字段")
		return
	}

	if isArticleExit(DB, title) {
		response.Fail(ctx, nil, "文章名重复")
		return
	}
	//更新文章表
	newArticle := model.Article{
		Model:       gorm.Model{},
		Title:       title,
		Body:        body,
		Describe:    describe,
		WriteTime:   writeTime,
		ReadingTime: readingTime,
		BackgroundImage: backgroundImage,
	}

	if err2 := DB.Create(&newArticle).Error; err2 != nil {
		response.Fail(ctx, nil, "文章创建失败")
		return
	}

	//更新tag表
	var tags []model.Tag
	for i := range tagArr {
		tag := model.Tag{
			Model:   gorm.Model{},
			TagName: tagArr[i].(string),
		}
		tags = append(tags, tag)
	}

	var arr2 []uint
	for i := range tags {
		id, ok := isTagExit(DB, tags[i].TagName)
		if !ok {
			if err3 := DB.Create(&tags[i]).Error; err3 != nil {
				response.Fail(ctx, nil, "tag创建失败")
				return
			}
			arr2 = append(arr2, tags[i].ID)
		} else {
			arr2 = append(arr2, id)
		}
	}

	//更新relation表
	var relations []model.Relation

	for i := range arr2 {
		relation := model.Relation{
			Model:     gorm.Model{},
			ArticleId: newArticle.ID,
			TagId:     arr2[i],
		}
		relations = append(relations, relation)
	}

	for i := range relations {
		if err4 := DB.Create(&relations[i]).Error; err4 != nil {
			response.Fail(ctx, nil, "relation创建失败")
			return
		}
	}

	response.Success(ctx, nil, "创建成功")
}

func GetAllTags(ctx *gin.Context)  {
	DB := common.GetDb()

	var tags []model.Tag
	err := DB.Find(&tags).Error

	if err != nil {
		response.Fail(ctx, gin.H{"err":err}, "获取失败")
		return
	}

	response.Success(ctx, gin.H{"tagNames":dto.ToTagsDto(tags)}, "获取成功")
}

func GetAllArticles(ctx *gin.Context)  {
	DB := common.GetDb()

	var articles []model.Article
	if err := DB.Find(&articles).Error; err != nil {
		response.Fail(ctx, gin.H{"err":err}, "获取失败")
		return
	}

	response.Success(ctx, gin.H{"articles": dto.ToArticlesDto(articles)}, "获取成功")
}

//获取标签对文章的关系 -- 只需调用一次即可 --> 返回值是key为tag_id,value为article_id数组的一个map转成的json
func GetRelationT(ctx *gin.Context)  {
	DB := common.GetDb()

	var relations []model.Relation
	err := DB.Find(&relations).Error

	if err != nil {
		response.Fail(ctx, gin.H{"err":err}, "获取失败")
		return
	}
	var arr [][]int
	for i := range relations {
		arr = append(arr, []int{int(relations[i].ArticleId), int(relations[i].TagId)})
	}

	res, err := json.Marshal(util.UnionRelation(arr))
	response.Success(ctx, gin.H{"res":string(res)}, "获取成功")
}

//获取标签对文章的关系 -- 只需调用一次即可 --> 返回值是key为article_id,value为tag_id数组的一个map转成的json
func GetRelationA(ctx *gin.Context)  {
	DB := common.GetDb()

	var relations []model.Relation
	err := DB.Find(&relations).Error

	if err != nil {
		response.Fail(ctx, gin.H{"err":err}, "获取失败")
		return
	}
	var arr [][]int
	for i := range relations {
		arr = append(arr, []int{int(relations[i].TagId), int(relations[i].ArticleId)})
	}

	res, err := json.Marshal(util.UnionRelation(arr))
	response.Success(ctx, gin.H{"res":string(res)}, "获取成功")
}

//通过标签选择的后端实现
func GetArticlesByRelation(ctx *gin.Context)  {
	//其实通过relation获取到的是一个与标签相关的文章ID的数组
	DB := common.GetDb()

	fjson := make(map[string]interface{})
	err1 := ctx.BindJSON(&fjson)
	if err1 != nil {
		response.Fail(ctx, gin.H{"err":err1}, "请求失败")
		return
	}

	arr := fjson["relation"].([]interface{})
	var articles []model.Article
	for i := range arr {
		newArticle := model.Article{
			Model:       gorm.Model{},
		}
		newArticle.ID = uint(arr[i].(float64))
		DB.Find(&newArticle)
		articles = append(articles, newArticle)
		newArticle = model.Article{}
	}

	response.Success(ctx, gin.H{"articles":dto.ToTitlesDto(articles)},"获取成功")
}

//前端通过路由来获取文章的方式
func GetArticleByTitle(ctx *gin.Context) {
	DB := common.GetDb()

	fjson := make(map[string]interface{})
	err1 := ctx.BindJSON(&fjson)
	if err1 != nil {
		response.Fail(ctx, gin.H{"err":err1}, "请求失败")
		return
	}

	var article model.Article
	err2 := DB.First(&article, "title = ?", fjson["title"].(string)).Error
	if err2 != nil {
		response.Fail(ctx, gin.H{"err": err2}, "请求失败")
		return
	}

	response.Success(ctx, gin.H{"article":dto.ToArticleDto(article)},"获取成功")
}

//文章的删除需要三个参数，文章号，文章对应的标签，关联
func DelArticle(ctx *gin.Context)  {
	DB := common.GetDb()

	//获取文章号,
	fjson := make(map[string]interface{})
	err1 := ctx.BindJSON(&fjson)
	if err1 != nil {
		response.Fail(ctx, gin.H{"err1":err1}, "请求失败")
		return
	}
	var article model.Article
	article.ID = uint(fjson["articleId"].(float64))

	//通过ID删除文章，因为设置了delete_at字段，故为软删除
	if err2 := DB.Delete(&article).Error; err2 != nil {
		response.Fail(ctx, gin.H{"err2":err2}, "请求失败")
		return
	}

	//删除关系表的字段
	if err3 := DB.Where("article_id = ?", uint(fjson["articleId"].(float64))).Delete(&model.Relation{}).Error; err3 != nil {
		response.Fail(ctx, gin.H{"err3":err3}, "请求失败")
		return
	}

	//是否需要删除空标签
	relation := fjson["relation"].([]interface{})
	var tag model.Tag
	for i := range relation {
		if isTagNeedDel(DB, uint(relation[i].(float64))) {
			tag.ID = uint(relation[i].(float64))
			if err4 := DB.Delete(&tag).Error; err4 != nil {
				response.Fail(ctx, gin.H{"err4":err4}, "请求失败")
				return
			}
		}
	}

	response.Success(ctx, nil, "删除成功")
}

//更新操作需要前后端结合进行 -- 由于将标签的部分进行了分离，只需关注文章的更新即可
func UpdateArticle(ctx *gin.Context)  {
	DB := common.GetDb()

	fjson := make(map[string]interface{})
	err1 := ctx.BindJSON(&fjson)
	if err1 != nil {
		response.Fail(ctx, gin.H{"err1":err1}, "请求失败")
		return
	}
	//对于文章内容
	newArticle := model.Article{}

	newArticle.ID = uint(fjson["articleId"].(float64))
	//直接更新
	if err2 := DB.First(&newArticle).Error; err2 != nil {
		response.Fail(ctx, gin.H{"err2": err2}, "文章不存在")
		return
	}

	newArticle.Title = fjson["title"].(string)
	newArticle.Body = fjson["body"].(string)
	newArticle.Describe = fjson["describe"].(string)
	newArticle.WriteTime = fjson["describe"].(string)
	newArticle.ReadingTime = fjson["readingTime"].(string)
	newArticle.BackgroundImage = fjson["backgroundImage"].(string)

	if err3 := DB.Save(&newArticle).Error; err3 != nil {
		response.Fail(ctx, gin.H{"err3": err3}, "更新失败")
	}
	response.Success(ctx, nil, "更新成功")
}

func CreateTagName(ctx *gin.Context)  {
	DB := common.GetDb()
	fjson := make(map[string]interface{})

	if err := ctx.BindJSON(&fjson); err != nil {
		response.Fail(ctx, gin.H{"err": err}, "绑定失败")
		return
	}

	tag := model.Tag{
		Model:   gorm.Model{},
		TagName: fjson["tagName"].(string),
	}

	relation := model.Relation{
		Model:     gorm.Model{},
		ArticleId: uint(fjson["articleId"].(float64)),
		TagId:     0,
	}
	//先判断是否是新标签
	id, ok := isTagNew(DB, tag.TagName)
	if ok {
		//是新标签
		if err2 := DB.Create(&tag).First(&tag).Error; err2 != nil {
			response.Fail(ctx, gin.H{"err2":err2}, "标签创建失败")
			return
		}
		relation.TagId = tag.ID
		if err3 := DB.Create(&relation).Error; err3 != nil {
			response.Fail(ctx, gin.H{"err3":err3}, "关系创建失败")
			return
		}
		response.Success(ctx, gin.H{"tagId":tag.ID}, "新建成功")
	} else {
		//不是新标签 -- 只需更新关系表
		relation.TagId = id
		if err := DB.Create(&relation).Error; err != nil {
			response.Fail(ctx, gin.H{"err": err}, "创建失败")
			return
		}
		response.Success(ctx, gin.H{"tagId":id}, "新建成功")
	}
}

func DelTagName(ctx *gin.Context) {
	DB := common.GetDb()
	fjson := make(map[string]interface{})
	if err := ctx.BindJSON(&fjson); err != nil {
		response.Fail(ctx, gin.H{"err":err}, "绑定失败")
	}

	tag := model.Tag{
		Model:   gorm.Model{},
		TagName: fjson["tagName"].(string),
	}
	tag.ID = uint(fjson["tagId"].(float64))

	relation := model.Relation{
		Model:     gorm.Model{},
		ArticleId: uint(fjson["articleId"].(float64)),
		TagId:     uint(fjson["tagId"].(float64)),
	}

	//删除关系
	if err := DB.Where("article_id = ? and tag_id = ?", relation.ArticleId, relation.TagId).Delete(relation).Error; err != nil {
		response.Fail(ctx, gin.H{"err":err}, "关系表删除失败")
	}

	//删除标签
	if isTagNeedDel(DB, tag.ID) {
		if err := DB.Delete(&tag).Error; err != nil {
			response.Fail(ctx, gin.H{"err": err}, "标签表删除失败")
		}
	}

	response.Success(ctx, nil, "删除成功")
}

//图片上传
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

//判断文章是否存在
func isArticleExit(db *gorm.DB, title string) bool {
	var article model.Article
	db.Where("title = ?", title).First(&article)
	if article.ID != 0 {
		return true
	}

	return false
}

//判断标签是否存在
func isTagExit(db *gorm.DB, tagName string) (uint, bool) {
	var tag model.Tag
	db.Where("tag_name = ?", tagName).First(&tag)
	if tag.ID != 0 {
		return tag.ID, true
	}

	return 0, false
}

//判断标签是否需要删除
func isTagNeedDel(db *gorm.DB, tagID uint) bool {
	var relation model.Relation
	db.Where("tag_id = ?", tagID).First(&relation)
	if relation.ID != 0 {
		return false
	}
	return true
}

//判断标签是否是新的
func isTagNew(db *gorm.DB, tagName string) (uint, bool) {
	var tag model.Tag
	db.Where("tag_name = ?", tagName).First(&tag)
	if tag.ID != 0 {
		return tag.ID, false
	}
	return 0, true
}
