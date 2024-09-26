package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/matthewyuh246/blogbackend/models"
	"github.com/matthewyuh246/blogbackend/usecase"
)

type IBlogController interface {
	CreatePost(c echo.Context) error
	GetAllPost(c echo.Context) error
	GetPostDetail(c echo.Context) error
	UpdatePost(c echo.Context) error
	UniquePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type blogController struct {
	bu usecase.IBlogUsecase
}

func NewBlogController(bu usecase.IBlogUsecase) IBlogController {
	return &blogController{bu}
}

func (bc *blogController) CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	blog := models.Blog{}
	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	blog.UserId = uint(userId.(float64))
	blogRes, err := bc.bu.CreatePost(blog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, blogRes)
}

func (bc *blogController) GetAllPost(c echo.Context) error {
	pageQuery := c.QueryParam("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 1 {
		page = 1
	}

	limit := 5

	blogs, currentPage, total, lastPage, err := bc.bu.GetAllPost(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blogs,
		"meta": map[string]interface{}{
			"total":     total,
			"page":      currentPage,
			"last_page": lastPage,
		},
	})
}

func (bc *blogController) GetPostDetail(c echo.Context) error {
	id := c.Param("blogId")
	blogId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid post ID",
		})
	}

	blog, err := bc.bu.GetPostDetail(uint(blogId))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Post not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blog,
	})
}

func (bc *blogController) UpdatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("blogId")
	blogId, _ := strconv.Atoi(id)
	blog := models.Blog{}
	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	blogRes, err := bc.bu.UpdatePost(blog, uint(userId.(float64)), uint(blogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, blogRes)
}

func (bc *blogController) UniquePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	blogs, err := bc.bu.GetBlogsByUserId(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, blogs)
}

func (bc *blogController) DeletePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("blogId")
	blogId, _ := strconv.Atoi(id)

	err := bc.bu.DeletePost(uint(userId.(float64)), uint(blogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
