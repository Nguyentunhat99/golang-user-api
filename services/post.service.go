package services

import "github.com/example/Nhat-golang-test/models"

type PostService interface {
	CreatePost(*models.CreatePostRequest) (*models.DBPost, error)
	UpdatePost(string, *models.UpdatePost) (*models.DBPost, error)
	DeletePost(string) error
	FindPostById(string) (*models.DBPost, error)
	FindPosts(page int, limit int) ([]*models.DBPost, error)
}
