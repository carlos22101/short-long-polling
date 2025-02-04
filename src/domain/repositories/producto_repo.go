package repositories

import "api-hexagonal-go/src/domain/entities"

type ProductoRepository interface {
	GetAll()([]entities.Producto,error)
	GetByID(id int) (*entities.Producto, error)
	Create(producto *entities.Producto) error
	Update(id int, producto *entities.Producto) error
	Delete(id int) error
}