package models

type Tramite struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Iniciales   string `json:"iniciales"`
	TerceraEdad string `json:"tercera_edad"`
}
