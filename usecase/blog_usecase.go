package usecase

import (
	"math"

	"github.com/matthewyuh246/blogbackend/models"
	repository "github.com/matthewyuh246/blogbackend/repositroy"
)

type IBlogUsecase interface {
	CreatePost(blog models.Blog) (models.BlogResponse, error)
	GetAllPost(page int, limit int) ([]models.Blog, int, int64, int, error)
	GetPostDetail(id uint) (*models.Blog, error)
	UpdatePost(blog models.Blog, userId uint, blogId uint) (models.BlogResponse, error)
	GetBlogsByUserId(userId uint) ([]models.Blog, error)
	DeletePost(userId uint, blogId uint) error
}

type blogUsecase struct {
	br repository.IBlogRepository
}

func NewBlogUsecase(br repository.IBlogRepository) IBlogUsecase {
	return &blogUsecase{br}
}

func (bu *blogUsecase) CreatePost(blog models.Blog) (models.BlogResponse, error) {
	if err := bu.br.CreatePost(&blog); err != nil {
		return models.BlogResponse{}, err
	}
	resBlog := models.BlogResponse{
		Id:        blog.Id,
		Title:     blog.Title,
		Desc:      blog.Desc,
		Image:     blog.Image,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
	return resBlog, nil
}

func (bu *blogUsecase) GetAllPost(page int, limit int) ([]models.Blog, int, int64, int, error) {
	offset := (page - 1) * limit

	blogs, total, err := bu.br.GetAllPost(offset, limit)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	return blogs, page, total, lastPage, nil
}

func (bu *blogUsecase) GetPostDetail(id uint) (*models.Blog, error) {
	return bu.br.GetPostByID(id)
}

func (bu *blogUsecase) UpdatePost(blog models.Blog, userId uint, blogId uint) (models.BlogResponse, error) {
	if err := bu.br.UpdatePost(&blog, userId, blogId); err != nil {
		return models.BlogResponse{}, err
	}
	resBlog := models.BlogResponse{
		Id:        blog.Id,
		Title:     blog.Title,
		Desc:      blog.Desc,
		Image:     blog.Image,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
	return resBlog, nil
}

func (bu *blogUsecase) GetBlogsByUserId(userId uint) ([]models.Blog, error) {
	return bu.br.FindByUserId(userId)
}

func (bu *blogUsecase) DeletePost(userId uint, blogId uint) error {
	if err := bu.br.DeletePost(userId, blogId); err != nil {
		return err
	}
	return nil
}
