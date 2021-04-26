package dto

import "demo2_backend/model"

type ArticleDto struct {
	Title string `json:"title"`
	Tag string `json:"tag"`
	Body string `json:"body"`
}

func ToArticleDto(article model.Article) ArticleDto {
	return ArticleDto{
		Title: article.Title,
		Tag:   article.Tag,
		Body:  article.Body,
	}
}

func ToArticlesDto(articles []model.Article) []ArticleDto {
	var arr []ArticleDto
	for i := range articles {
		arr = append(arr, ToArticleDto(articles[i]))
	}
	return arr
}

