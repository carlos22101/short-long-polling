package applications

import (
	"api-hexagonal-go/src/domain/entities"
	"api-hexagonal-go/src/domain/repositories"
)

type ProductoUseCase struct {
	Repo repositories.ProductoRepository
}
//contructor no olvidar e investigar
func NewProductoUseCase(repo repositories.ProductoRepository) *ProductoUseCase {
	return &ProductoUseCase{Repo: repo}
}

func (p *ProductoUseCase) GetAllProductos() ([]entities.Producto, error){
	return p.Repo.GetAll()
}
func (p *ProductoUseCase) GetProductoByID(id int) (*entities.Producto, error) {
	return p.Repo.GetByID(id)
}

func (p *ProductoUseCase) CreateProducto(producto *entities.Producto) error {
	return p.Repo.Create(producto)
}

func (p *ProductoUseCase) UpdateProducto(id int, producto *entities.Producto) error {
	return p.Repo.Update(id, producto)
}

func (p *ProductoUseCase) DeleteProducto(id int) error {
	return p.Repo.Delete(id)
}