package usecase

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"time"

	repository "github.com/matthewyuh246/blogbackend/repositroy"
)

type IImageUsecase interface {
	UploadFile(files []*multipart.FileHeader) (string, error)
}

type imageUsecase struct {
	ir repository.IImageRepository
}

func NewImageUsecase(ir repository.IImageRepository) IImageUsecase {
	return &imageUsecase{ir}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (iu *imageUsecase) UploadFile(files []*multipart.FileHeader) (string, error) {
	rand.Seed(time.Now().UnixNano())

	file := files[0]
	fileName := randLetter(5) + "-" + file.Filename

	err := iu.ir.SaveFile(file, fileName)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return "http://localhost:8080/api/blog/uploads/" + fileName, nil
}
