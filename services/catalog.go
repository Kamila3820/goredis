package services

type Product struct {
	ID       int
	Name     string
	Quantity int
}

type ICatalogService interface {
	GetProducts() ([]Product, error)
}
