package handlers

import (
	"goredis/services"

	"github.com/gofiber/fiber/v3"
)

type catalogHandler struct {
	catalogService services.ICatalogService
}

func NewCatalogHandler(catalogService services.ICatalogService) ICatalogHandler {
	return &catalogHandler{
		catalogService: catalogService,
	}
}

func (h *catalogHandler) GetProducts(c fiber.Ctx) error {
	products, err := h.catalogService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	return c.JSON(response)
}
