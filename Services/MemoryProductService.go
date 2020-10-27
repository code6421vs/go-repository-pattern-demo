package services

import repositorys "awesomeProject4/Repositorys"

func NewMemoryProductService(repository repositorys.ProductRepository) ProductService {
	return ProductService{Repository: repository}
}

func (s ProductService) AddProduct(p repositorys.Product) {
	s.Repository.Add(p)
}
