package dto

import "demo2_backend/model"

type ArticleDto struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Describe string `json:"describe"`
	WriteTime string `json:"writeTime"`
	ReadingTime string `json:"readingTime"`
	BackgroundImage string `json:"backgroundImage"`
}

type TitleDto struct {
	Title string `json:"title"`
	Describe string `json:"describe"`
	WriteTime string `json:"writeTime"`
	ReadingTime string `json:"readingTime"`
	BackgroundImage string `json:"backgroundImage"`
}

type UploadArticles struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Describe string `json:"describe"`
	WriteTime string `json:"writeTime"`
	ReadingTime string `json:"readingTime"`
	BackgroundImage string `json:"backgroundImage"`
}

func ToTitleDto(article model.Article) TitleDto {
	return TitleDto{
		Title:       article.Title,
		Describe: 	 article.Describe,
		WriteTime:   article.WriteTime,
		ReadingTime: article.ReadingTime,
		BackgroundImage: article.BackgroundImage,
	}
}

func ToTitlesDto(titles []model.Article) []TitleDto {
	var arr []TitleDto
	for i := range titles {
		arr = append(arr, ToTitleDto(titles[i]))
	}
	return arr
}

func ToArticleDto(article model.Article) ArticleDto {
	return ArticleDto{
		Id: article.ID,
		Title: article.Title,
		Body:  article.Body,
		Describe: article.Describe,
		WriteTime: article.WriteTime,
		ReadingTime: article.ReadingTime,
		BackgroundImage: article.BackgroundImage,
	}
}

func ToArticlesDto(articles []model.Article) []ArticleDto {
	var res []ArticleDto
	for i := range articles {
		res = append(res, ToArticleDto(articles[i]))
	}
	return res
}



