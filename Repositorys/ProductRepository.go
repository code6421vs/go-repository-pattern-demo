package repositorys

type Product struct {
	Id int
	Name string
	Price int
}

type ProductRepository interface {
	Add(p Product)
	Remove(p Product)
	Get() []Product
}

