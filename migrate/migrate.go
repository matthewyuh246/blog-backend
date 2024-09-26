package main

import (
	"fmt"

	"github.com/matthewyuh246/blogbackend/db"
	"github.com/matthewyuh246/blogbackend/models"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&models.User{}, &models.Blog{})
}
