package repositories

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return &productRepositoryRedis{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *productRepositoryRedis) GetProducts() ([]product, error) {
	return nil, nil
}
