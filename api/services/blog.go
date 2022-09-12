
package services

import (
	"blog/api/repositories"
	"blog/models"
)

//PostService PostService struct
type PostService struct {
	repositories repositories.PostRepository
}

//NewPostService : returns the PostService struct instance
func NewPostService(r repositories.PostRepository) PostService {
	return PostService{
		repositories: r,
	}
}

//Save -> calls post repository save method
func (p PostService) Save(post models.Post) error {
	return p.repositories.Save(post)
}

//FindAll -> calls post repo find all method
func (p PostService) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	return p.repositories.FindAll(post, keyword)
}

// Update -> calls postrepo update method
func (p PostService) Update(post models.Post) error {
	return p.repositories.Update(post)
}

// Delete -> calls post repo delete method
func (p PostService) Delete(id int64) error {
	var post models.Post
	post.ID = id
	return p.repositories.Delete(post)
}

// Find -> calls post repo find method
func (p PostService) Find(post models.Post) (models.Post, error) {
	return p.repositories.Find(post)
}
