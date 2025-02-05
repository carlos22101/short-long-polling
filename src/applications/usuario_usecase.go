package applications

import (
	"api-hexagonal-go/src/domain/entities"
	"api-hexagonal-go/src/domain/repositories"
	"api-hexagonal-go/src/infraestructure/controllers/notifier"
	// ✅ Importar el notifier sin ciclos
)

type UsuarioUseCase struct {
	Repo     repositories.Usuariorepository
	Notifier *notifier.Notifier // ✅ Usamos el Notifier
}

func NewUsuarioUseCase(repo repositories.Usuariorepository, notifier *notifier.Notifier) *UsuarioUseCase {
	return &UsuarioUseCase{Repo: repo, Notifier: notifier}
}

func (u *UsuarioUseCase) GetAllUsuarios() ([]entities.Usuario, error) {
	return u.Repo.GetAll()
}

func (u *UsuarioUseCase) GetUsuarioByID(id int) (*entities.Usuario, error) {
	return u.Repo.GetByID(id)
}

func (u *UsuarioUseCase) CreateUsuario(usuario *entities.Usuario) error {
	err := u.Repo.Create(usuario)
	if err == nil {
		u.Notifier.NotifyChanges() // 🔥 Notificar cambios
	}
	return err
}

func (u *UsuarioUseCase) UpdateUsuario(id int, usuario *entities.Usuario) error {
	err := u.Repo.Update(id, usuario)
	if err == nil {
		u.Notifier.NotifyChanges() // 🔥 Notificar cambios
	}
	return err
}

func (u *UsuarioUseCase) DeleteUsuario(id int) error {
	err := u.Repo.Delete(id)
	if err == nil {
		u.Notifier.NotifyChanges() // 🔥 Notificar cambios
	}
	return err
}
