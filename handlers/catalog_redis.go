package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
)

type catalogHandlerRedis struct {
	productService services.ICatalogService
	redisClient    *redis.Client
}

func NewCatalogHandlerRedis(productService services.ICatalogService, redisClient *redis.Client) ICatalogHandler {
	return &catalogHandlerRedis{
		productService: productService,
		redisClient:    redisClient,
	}
}

func (h *catalogHandlerRedis) GetProducts(c fiber.Ctx) error {
	key := "handler::GetProducts"

	// Redis GET
	if responseJSON, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("Redis Handler")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJSON)
	}

	// Services
	products, err := h.productService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// Redis SET
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("Services")

	return c.JSON(response)
}
