package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repositories.IProductRepository
	redisClient *redis.Client
}

func NewCatalogSreviceRedis(productRepo repositories.IProductRepository, redisClient *redis.Client) ICatalogService {
	return &catalogServiceRedis{
		productRepo: productRepo,
		redisClient: redisClient,
	}
}

func (s *catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "service::GetProducts"

	// Redis GET
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		err = json.Unmarshal([]byte(productJson), &products)
		if err == nil {
			fmt.Println("Redis Service")
			return products, nil
		}
	}

	// Repositories
	productsRepo, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productsRepo {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	// Redis SET
	if productByte, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(productByte), time.Second*10)
	}

	fmt.Println("Database")

	return products, nil
}
