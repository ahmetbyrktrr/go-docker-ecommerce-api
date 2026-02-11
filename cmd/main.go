package main

import (
	"database/sql"
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres",
		"host=db user=postgres password=postgres dbname=ecommerce sslmode=disable")

	r := gin.Default()

	catRepo := &repository.CategoryRepository{DB: db}
	catService := &service.CategoryService{Repo: catRepo}
	catHandler := &handler.CategoryHandler{Service: catService}
	catHandler.Register(r)

	prodRepo := &repository.ProductRepository{DB: db}
	prodService := &service.ProductService{Repo: prodRepo}
	prodHandler := &handler.ProductHandler{Service: prodService}
	prodHandler.Register(r)

	r.Run(":8080")
}
