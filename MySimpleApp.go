package main

import (
	services "awesomeProject4/Services"
	"fmt"
	"awesomeProject4/Repositorys"
)


func work(s services.ProductService) {
	s.AddProduct(repositorys.Product{Id: 1, Name: "test", Price: 12})
}

func main() {
	s := NewProductService()
	work(s)
	fmt.Println("Done")
}
