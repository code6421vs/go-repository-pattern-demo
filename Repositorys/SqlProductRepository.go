package repositorys

type SqlProductRepository struct {
	ConnectionString string
}

func NewSqlProductRepository() SqlProductRepository {
	return SqlProductRepository{""}
}

func (s SqlProductRepository) Add(p Product) {
}

func (s SqlProductRepository) Remove(p Product) {

}

func (s SqlProductRepository) Get() []Product {
	return make([]Product, 20)
}
