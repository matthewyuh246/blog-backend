package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/matthewyuh246/blogbackend/models"
	repository "github.com/matthewyuh246/blogbackend/repositroy"
	"github.com/matthewyuh246/blogbackend/validator"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user models.User) (models.UserResponse, error)
	Login(user models.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user models.User) (models.UserResponse, error) {
	if err := uu.uv.SignUpUserValidate(user); err != nil {
		return models.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return models.UserResponse{}, err
	}
	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}
	resUser := models.UserResponse{
		Id:    newUser.Id,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user models.User) (string, error) {
	if err := uu.uv.LoginUserValidate(user); err != nil {
		return "", err
	}
	storedUser := models.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.Id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
