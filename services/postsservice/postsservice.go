package postsservice

import (
	"github.com/aubm/blog/models"
	"github.com/aubm/blog/services/database"
)

func GetPosts() []models.Post {
	db := database.GetDatabaseLayer()
	var posts []models.Post
	db.Find(&posts)
	return posts
}

func GetOnePost(postId int64) models.Post {
	db := database.GetDatabaseLayer()
	var post models.Post
	db.Where("id = ?", postId).Find(&post)
	return post
}

func SavePost(post *models.Post) {
	db := database.GetDatabaseLayer()
	if db.NewRecord(post) {
		db.Create(post)
	} else {
		db.Save(post)
	}
}
