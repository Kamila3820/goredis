package main

import (
	"fmt"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogSreviceRedis(productRepo, redisClient)

	products, err := productService.GetProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result: ")
	fmt.Println(products)
	// app := fiber.New()

	// app.Get("/hello", func(c fiber.Ctx) error {
	// 	// time.Sleep(time.Millisecond * 10)
	// 	return c.SendString("Hello")
	// })

	// app.Listen(":8000")
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
