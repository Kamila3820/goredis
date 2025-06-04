package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) IProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return &productRepositoryRedis{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *productRepositoryRedis) GetProducts() (products []product, err error) {
	key := "repository::GetProducts"

	// Redis get
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("Redis Repository")
			return products, nil
		}
	}

	// Database
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Redis set
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database")

	return nil, nil
}
