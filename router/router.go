package router

import (
	// "net/http"
	// "os"

	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/matthewyuh246/blogbackend/controller"
)

func NewRouter(uc controller.IUserController, bc controller.IBlogController, ic controller.IImageController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/api",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
	}))
	t := e.Group("/api")
	t.POST("/signup", uc.SignUp)
	t.POST("/login", uc.Login)
	t.POST("/logout", uc.Logout)
	t.GET("/csrf", uc.CsrfToken)
	s := t.Group("/blog")
	s.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	s.POST("/post", bc.CreatePost)
	s.PUT("/updatepost/:blogId", bc.UpdatePost)
	s.GET("/uniquepost", bc.UniquePost)
	s.DELETE("/deletepost/:blogId", bc.DeletePost)
	s.POST("/upload-image", ic.Upload)
	s.Static("/uploads", "./uploads")
	u := s.Group("/getpost")
	u.GET("/all", bc.GetAllPost)
	u.GET("/:blogId", bc.GetPostDetail)
	return e
}
