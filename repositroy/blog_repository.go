package repository

import (
	"fmt"

	"github.com/matthewyuh246/blogbackend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBlogRepository interface {
	CreatePost(blog *models.Blog) error
	GetAllPost(offset int, limit int) ([]models.Blog, int64, error)
	GetPostByID(id uint) (*models.Blog, error)
	UpdatePost(blog *models.Blog, userId uint, blogId uint) error
	FindByUserId(userId uint) ([]models.Blog, error)
	DeletePost(userId uint, blogId uint) error
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) IBlogRepository {
	return &blogRepository{db}
}

func (br *blogRepository) CreatePost(blog *models.Blog) error {
	if err := br.db.Create(blog).Error; err != nil {
		return err
	}
	return nil
}

func (br *blogRepository) GetAllPost(offset int, limit int) ([]models.Blog, int64, error) {
	var blogs []models.Blog
	var total int64

	if err := br.db.Preload("User").Offset(offset).Limit(limit).Find(&blogs).Error; err != nil {
		return nil, 0, err
	}

	br.db.Model(&models.Blog{}).Count(&total)

	return blogs, total, nil
}

func (br *blogRepository) GetPostByID(id uint) (*models.Blog, error) {
	var blog models.Blog
	if err := br.db.Where("id = ?", id).Preload("User").First(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil

}

func (br *blogRepository) UpdatePost(blog *models.Blog, userId uint, blogId uint) error {
	result := br.db.Model(blog).
		Clauses(clause.Returning{}).
		Where("id=? AND user_id=?", blogId, userId).
		Updates(map[string]interface{}{
			"title": blog.Title,
			"desc":  blog.Desc,
			"image": blog.Image,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (br *blogRepository) FindByUserId(userId uint) ([]models.Blog, error) {
	var blog []models.Blog
	result := br.db.Model(&blog).Where("User_id = ?", userId).Preload("User").Find(&blog)
	return blog, result.Error
}

func (br *blogRepository) DeletePost(userId uint, blogId uint) error {
	result := br.db.Where("id=? AND user_id=?", blogId, userId).Delete(&models.Blog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
