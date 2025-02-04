package repositories

import (
    "api-hexagonal-go/src/domain/entities"
    "database/sql"
)

type UsuarioRepository struct {
    DB *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
    return &UsuarioRepository{DB: db}
}


func (r *UsuarioRepository) GetAll() ([]entities.Usuario, error) {
    rows, err := r.DB.Query("SELECT id, nombre, email, edad FROM usuarios")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var usuarios []entities.Usuario
    for rows.Next() {
        var u entities.Usuario
        if err := rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad); err != nil {
            return nil, err
        }
        usuarios = append(usuarios, u)
    }
    return usuarios, nil
}


func (r *UsuarioRepository) Create(usuario *entities.Usuario) error {
    _, err := r.DB.Exec("INSERT INTO usuarios (nombre, email, edad) VALUES (?, ?, ?)", usuario.Nombre, usuario.Email, usuario.Edad)
    return err
}

func (r *UsuarioRepository) GetByID(id int) (*entities.Usuario, error) {
    row := r.DB.QueryRow("SELECT id, nombre, email, edad FROM usuarios WHERE id = ?", id)
    var u entities.Usuario
    if err := row.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &u, nil
}


func (r *UsuarioRepository) Update(id int, usuario *entities.Usuario) error {
    _, err := r.DB.Exec("UPDATE usuarios SET nombre = ?, email = ?, edad = ? WHERE id = ?", usuario.Nombre, usuario.Email, usuario.Edad, id)
    return err
}


func (r *UsuarioRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM usuarios WHERE id = ?", id)
    return err
}