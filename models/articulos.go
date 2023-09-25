package models

import "gorm.io/gorm"

type Articulos struct {
	gorm.Model
	ID          int64  `gorm:"not null; unique_index"`
	Idcategoria int64  `gorm:"not null; unique_index; default:0"`
	Codigo      string `gorm:"not null; unique_index"`
	Articulo    string `gorm:"not null"`
	Descripcion string
	Precio      float64 `gorm:"default:0"`
}
