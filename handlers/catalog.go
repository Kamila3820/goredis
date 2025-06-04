package handlers

import "github.com/gofiber/fiber/v3"

type ICatalogHandler interface {
	GetProducts(c fiber.Ctx) error
}
