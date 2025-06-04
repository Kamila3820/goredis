package main

import (
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogSreviceRedis(productRepo, redisClient)
	productHandler := handlers.NewCatalogHandler(productService)

	app := fiber.New()
	app.Get("/products", productHandler.GetProducts)
	app.Listen(":8000")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/goredis")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
