package applications

import (
	"api-hexagonal-go/src/domain/entities"
	"api-hexagonal-go/src/domain/repositories"
)

type UsuarioUseCase struct {
	Repo repositories.Usuariorepository
}

func NewUsuariouseCase(repo repositories.Usuariorepository) *UsuarioUseCase{
	return &UsuarioUseCase{Repo: repo}
}

func(u*UsuarioUseCase) GetAllUsuarios()([]entities.Usuario, error){
	return u.Repo.GetAll()
}
func (u *UsuarioUseCase) GetUsuarioByID(id int) (*entities.Usuario, error) {
	return u.Repo.GetByID(id)
}

func (u *UsuarioUseCase) CreateUsuario(usuario *entities.Usuario) error {
	return u.Repo.Create(usuario)
}

func (u *UsuarioUseCase) UpdateUsuario(id int, usuario *entities.Usuario) error {
	return u.Repo.Update(id, usuario)
}

func (u *UsuarioUseCase) DeleteUsuario(id int) error {
	return u.Repo.Delete(id)
}