package serializer

import "ginexample/models"

// Article序列化器
type Tag struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
}


func BuildTag(item models.Tag) Tag {
	return Tag{
		ID: item.ID,
		Name: item.Name,
		CreatedBy: item.CreatedBy,
		CreatedAt: item.CreatedAt.Unix(),
	}
}


// BuildVideos 序列化视频列表
func BuildTags(items []models.Tag) (tags []Tag) {
	for _, item := range items {
		tag := BuildTag(item)
		tags = append(tags, tag)
	}
	return tags
}

