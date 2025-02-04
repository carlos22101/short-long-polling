package entities

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"correo"`
	Edad   int    `json:"edad"`
}
