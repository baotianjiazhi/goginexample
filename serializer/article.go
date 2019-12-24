package serializer

import "ginexample/models"

// Article序列化器
type Article struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}


func BuildArticle(item models.Article) Article {
	return Article{
		ID: item.ID,
		Title: item.Title,
		Content: item.Content,
		CreatedAt: item.CreatedAt.Unix(),
	}
}


// BuildVideos 序列化视频列表
func BuildArticles(items []models.Article) (videos []Article) {
	for _, item := range items {
		video := BuildArticle(item)
		videos = append(videos, video)
	}
	return videos
}

