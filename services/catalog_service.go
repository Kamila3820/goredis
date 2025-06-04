package services

import "goredis/repositories"

type catalogService struct {
	productRepo repositories.IProductRepository
}

func NewCatalogService(productRepo repositories.IProductRepository) ICatalogService {
	return &catalogService{
		productRepo: productRepo,
	}
}

func (s *catalogService) GetProducts() ([]Product, error) {
	products, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	productService := make([]Product, 0)
	for _, p := range products {
		productService = append(productService, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return productService, nil
}
