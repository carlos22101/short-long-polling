package repositories

import "api-hexagonal-go/src/domain/entities"

type Usuariorepository interface {
	GetAll()([]entities.Usuario, error)
	GetByID(id int) (*entities.Usuario, error)
	Create(usuario *entities.Usuario) error
	Update(id int, usuario *entities.Usuario) error
	Delete(id int) error
}