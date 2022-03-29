package models

type Persona struct {
	ID        int64 `json "id" gorm:"primary_key;auto_increment"`
	Nombre    int64 `json:"nombre"`
	Apellido  int64 `json:"apellido"`
	Direccion int64 `json:"direccion"`
	Telefono  int64 `json:"telefono"`
}
