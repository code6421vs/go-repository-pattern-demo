package main

import (
	repositorys "awesomeProject4/Repositorys"
	services "awesomeProject4/Services"
)

func NewProductService() services.ProductService {
	resp := repositorys.NewSqlProductRepository()
	return services.NewMemoryProductService(resp)
}
