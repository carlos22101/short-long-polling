package repositories

import (
    "api-hexagonal-go/src/domain/entities"
    "database/sql"
)

type ProductoRepository struct {
    DB *sql.DB
}

func NewProductoRepository(db *sql.DB) *ProductoRepository {
    return &ProductoRepository{DB: db}
}

func (r *ProductoRepository) GetAll() ([]entities.Producto, error) {
    rows, err := r.DB.Query("SELECT id, nombre, descripcion, precio FROM productos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var productos []entities.Producto
    for rows.Next() {
        var p entities.Producto
        if err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio); err != nil {
            return nil, err
        }
        productos = append(productos, p)
    }
    return productos, nil
}

func (r *ProductoRepository) GetByID(id int) (*entities.Producto, error) {
    row := r.DB.QueryRow("SELECT id, nombre, descripcion, precio FROM productos WHERE id = ?", id)
    var p entities.Producto
    if err := row.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &p, nil
}

func (r *ProductoRepository) Create(producto *entities.Producto) error {
    _, err := r.DB.Exec("INSERT INTO productos (nombre, descripcion, precio) VALUES (?, ?, ?)", producto.Nombre, producto.Descripcion, producto.Precio)
    return err
}

func (r *ProductoRepository) Update(id int, producto *entities.Producto) error {
    _, err := r.DB.Exec("UPDATE productos SET nombre = ?, descripcion = ?, precio = ? WHERE id = ?", producto.Nombre, producto.Descripcion, producto.Precio, id)
    return err
}

func (r *ProductoRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM productos WHERE id = ?", id)
    return err
}