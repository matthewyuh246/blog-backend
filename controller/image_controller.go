package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/matthewyuh246/blogbackend/usecase"
)

type IImageController interface {
	Upload(c echo.Context) error
}

type imageController struct {
	iu usecase.IImageUsecase
}

func NewImageController(iu usecase.IImageUsecase) IImageController {
	return &imageController{iu}
}

func (ic *imageController) Upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil || form == nil {
		return c.JSON(http.StatusBadRequest, "Invalid form data")
	}

	files, ok := form.File["image"]
	if !ok || len(files) == 0 {
		return c.JSON(http.StatusBadRequest, "No files uploaded")
	}

	url, err := ic.iu.UploadFile(files)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"url": url,
	})
}
