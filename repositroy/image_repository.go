package repository

import (
	"io"
	"mime/multipart"
	"os"
)

type IImageRepository interface {
	SaveFile(file *multipart.FileHeader, filename string) error
}

type imageRepository struct {
	uploadPath string
}

func NewImageRepository(uploadPath string) IImageRepository {
	return &imageRepository{uploadPath}
}

func (ir *imageRepository) SaveFile(file *multipart.FileHeader, filename string) error {
	dst := ir.uploadPath + filename
	if err := os.MkdirAll(ir.uploadPath, os.ModePerm); err != nil {
		return err
	}
	return saveMultipartFile(file, dst)
}

func saveMultipartFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
