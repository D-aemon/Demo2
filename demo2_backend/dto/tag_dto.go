package dto

import "demo2_backend/model"

type TagDto struct {
	Id uint `json:"id"`
	TagName string `json:"TagName"`
}

func ToTagDto(tag model.Tag) TagDto {
	return TagDto{
		Id: tag.ID,
		TagName: tag.TagName,
	}
}

func ToTagsDto(tags []model.Tag) []TagDto {
	var arr []TagDto

	for i := range tags {
		arr = append(arr, ToTagDto(tags[i]))
	}

	return arr
}